// ref: https://hacktam.kr/etclec/52
package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
)

func CreateAccessToken(userId uint64) (string, error) {
	log.Info("request create access token for userId=", userId)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() // access 토큰의 유효시간은 발급 순간부터 10분
	claims["user_id"] = userId

	encToken, err := token.SignedString([]byte(config.Config.ApiGw.AccessTokenSecret))
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Info("Token Decoding Success -> userId=", userId)
	return encToken, nil
}

func VerifyAccessToken(t string) (uint64, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// 파라미터로 받은 토큰이 HMAC 알고리즘이 맞는지 확인
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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
		return 0, fmt.Errorf("token not invalid")
	}

	userId := uint64(claims["user_id"].(float64)) // TODO: 이렇게 타입변환해도 괜찮나? 숫자가 커졌을 때도 문제없는지 확인필요..
	log.Info("token verifed userId=", userId)
	return userId, nil
}
