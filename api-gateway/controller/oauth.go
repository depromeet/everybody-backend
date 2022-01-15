package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/depromeet/everybody-backend/api-gateway/model"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	googeAuthUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	kakaoAuthUrlAPI = "https://kapi.kakao.com/v2/user/me"
	appleAuthUrlAPI
)

type AuthRequest struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
}

func getKakaoUserInfo(accessToken string) ([]byte, error) {
	req, err := http.NewRequest("POST", kakaoAuthUrlAPI, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	return userInfo, nil
}

func getGoogleUserInfo(accessToken string) ([]byte, error) {
	resp, err := http.Get(googeAuthUrlAPI + accessToken)
	// accessToken이 만료되거나 invalid 할 때 error catch가 안되는 거 같은데 확인 필요
	if err != nil {
		log.Error(err.Error())
		return nil, errors.WithStack(err)
	}

	userInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	return userInfo, nil
}

func OauthLogin(c echo.Context) error {
	authRequest := &AuthRequest{}
	if err := c.Bind(authRequest); err != nil {
		log.Error("json body parse error...", "\nerr=", err)
		return c.String(400, "잘못된 요청입니다")
	}

	ua, err := model.GetUserAuth(authRequest.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("user not found... req=", authRequest)
			return c.String(401, "user not found...")
		}
		log.Error(err)
		return c.String(500, "알 수 없는 오류가 발생했습니다")
	}

	if authRequest.Password != ua.Password {
		log.Info("password unmatched... req=", authRequest)
		return c.String(401, "password unmatched...")
	}

	// 소셜 info 불러오기
	accessToken := c.Get("oauthtoken").(string)
	var userInfo []byte
	switch authRequest.Kind {
	case "KAKAO":
		kakaoInfo, err := getKakaoUserInfo(accessToken)
		if err != nil {
			log.Error(err.Error())
			return c.String(500, err.Error())
		}
		userInfo = kakaoInfo
	case "GOOGLE":
		googleInfo, err := getGoogleUserInfo(accessToken)
		if err != nil {
			log.Error(err.Error())
			return c.String(500, err.Error())
		}
		userInfo = googleInfo
	}

	// TODO: 구조체로 만들기
	var socialUser map[string]interface{}
	if err := json.Unmarshal(userInfo, &socialUser); err != nil {
		log.Error("json body parse error...", "\nerr=", err)
		return c.String(500, "잘못된 요청입니다")
	}
	log.Info("소셜 로그인으로 user 정보 획득", string(userInfo))

	// kakao는 {"msg": , "code": ""} 토큰 invalid하면 이렇게 응답 옴, google은 {"error":}
	if socialUser["code"] != nil || socialUser["error"] != nil {
		log.Error("social token invalid")
		return c.String(401, "invalid token")
	}

	if socialUser["id"] == nil {
		log.Error("socialUser 값이 존재 하지 않음...")
		return c.String(401, "invalid token")
	}

	// google은 id 값이 string
	socialId, ok := socialUser["id"].(string)
	if !ok {
		// kakao는 id 값이 float64
		socialId = strconv.Itoa(int(socialUser["id"].(float64)))
	}

	// 소셜로그인으로 id 값 획득하고 소셜 id로 UserAuth 테이블 탐색 후 존재하지 않으면 해당 userId에 social_id 저장
	userAuth, err := model.GetUserAuthBySocialId(socialId)
	if err != nil {
		// does not exist
		if err == sql.ErrNoRows {
			err := model.SetUserAuthWithSocialId(authRequest.UserId, socialId)
			if err != nil {
				return c.String(500, "알 수 없는 오류가 발생했습니다")
			}
		} else {
			return c.String(500, "알 수 없는 오류가 발생했습니다")
		}
	}

	log.Info("소셜id로 유저 조회 완료")
	// UserAuth의 user_id, social_id 값과 request body로 요청 오는 user_id 값이 다르면 잘못된 요청
	// userAuth nil 검사하는 이유는 최초로 소셜 로그인 하는 경우 userAuth 값이 없기 때문
	if userAuth != nil {
		if authRequest.UserId != userAuth.UserId {
			log.Error("소셜 인증을 요청한 유저 -> ", authRequest.UserId, " 기존 유저 -> ", userAuth.UserId)
			return c.String(401, "인증 실패")
		}
	}

	// UserAuth의 userId 값으로 앱 내에서 사용할 토큰 생성
	token, err := util.CreateAccessToken(authRequest.UserId)
	if err != nil {
		return c.String(500, "알 수 없는 오류가 발생했습니다.")
	}

	res := map[string]string{
		"access_token": token,
	}
	log.Info("Login ok... userId=", authRequest.UserId)

	return c.JSON(200, res)
}
