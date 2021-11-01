// ref: https://hacktam.kr/etclec/52
package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
)

//TODO: jwt 인증 절차 전체를 미들웨어?로 빼는 방안 고려 필요

func CreateAccessToken(userId int) (string, error) {
	log.Info("request create access token for userId=", userId)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Duration(config.Config.ApiGw.AccessTokenExpireTimeMin * 60000000000)).Unix() // 60000000000ns = 1min = time.Minute
	claims["user_id"] = strconv.Itoa(userId)

	encToken, err := token.SignedString([]byte(config.Config.ApiGw.AccessTokenSecret))
	if err != nil {
		log.Error(err)
		return "", err
	}
	encToken = "Bearer " + encToken

	log.Info("Token create Success -> userId=", userId)
	return encToken, nil
}

func VerifyAccessToken(t string) (int, error) {
	tokenTypePrefix := "Bearer "
	if !strings.HasPrefix(t, tokenTypePrefix) {
		return 0, fmt.Errorf("token type invalid")
	}
	t = t[len(tokenTypePrefix):]

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// 파라미터로 받은 토큰이 HMAC 알고리즘이 맞는지 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.ApiGw.AccessTokenSecret), nil
	})
	if err != nil {
		log.Error("token parse error... err=", err) // 토큰 유효시간이 만료된 경우 여기서 에러발생 - "Token is expired"
		return 0, err
	}

	// 토큰 유효성 확인
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("token not valid")
	}

	userId, err := strconv.Atoi(claims["user_id"].(string))
	if err != nil {
		return 0, fmt.Errorf("token not valid")
	}

	log.Info("token verified userId=", userId)
	return userId, nil
}
