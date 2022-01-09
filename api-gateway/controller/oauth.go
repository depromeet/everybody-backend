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

func getKakaoUserInfo(accessToken string) ([]byte, error) {
	req, err := http.NewRequest("POST", kakaoAuthUrlAPI, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	userInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return userInfo, nil
}

func getGoogleUserInfo(authURL, accessToken string) ([]byte, error) {
	resp, err := http.Get(authURL + accessToken)
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

func KakaoLogin(c echo.Context) error {
	accessToken := c.Get("oauthtoken").(string)
	userInfo, err := getKakaoUserInfo(accessToken)
	if err != nil {
		log.Error(err.Error())
		return c.String(500, err.Error())
	}

	var user map[string]int
	if err := c.Bind(&user); err != nil {
		return c.String(400, err.Error())
	}

	// TODO: 구조체로 만들기
	var socialUser map[string]interface{}
	if err := json.Unmarshal(userInfo, &socialUser); err != nil {
		return c.String(500, err.Error())
	}

	log.Info("소셜 로그인으로 user 정보 획득", string(userInfo))

	// kakao는 {"msg": , "code": ""} 토큰 invalid하면 이렇게 응답 옴
	if socialUser["code"] != nil {
		log.Error("token invalid")
		return c.String(401, "invalid token")
	}

	if socialUser["id"] == nil {
		log.Error("socialUser 값이 존재 하지 않음...")
		return c.String(401, "invalid token")
	}

	// 소셜로그인으로 id 값 획득하고 소셜 id로 UserAuth 테이블 탐색 후 존재하지 않으면 해당 userId에 social_id 저장
	socialId := strconv.Itoa(int(socialUser["id"].(float64)))
	userAuth, err := model.GetUserAuthBySocialId(socialId)
	if err != nil {
		// does not exist
		if err == sql.ErrNoRows {
			err := model.SetUserAuthWithSocialId(user["user_id"], socialId)
			if err != nil {
				return c.String(500, err.Error())
			}
		} else {
			return c.String(500, err.Error())
		}
	}

	// UserAuth의 user_id, social_id 값과 request body로 요청 오는 user_id 값이 다르면 잘못된 요청
	if userAuth != nil {
		if user["user_id"] != userAuth.UserId {
			return c.String(401, "이미 가입한 유저입니다.")
		}
	}
	// UserAuth의 userId 값으로 앱 내에서 사용할 토큰 생성
	token, err := util.CreateAccessToken(user["user_id"])
	if err != nil {
		return c.String(500, err.Error())
	}

	res := map[string]string{
		"access_token": token,
	}

	return c.JSON(200, res)
}

func GoogleLogin(c echo.Context) error {
	accessToken := c.Get("oauthtoken").(string)
	userInfo, err := getGoogleUserInfo(googeAuthUrlAPI, accessToken)
	if err != nil {
		log.Error(err.Error())
		return c.String(500, err.Error())
	}

	var user map[string]int
	if err := c.Bind(&user); err != nil {
		return c.String(400, err.Error())
	}

	// TODO: 구조체로 만들기
	var socialUser map[string]interface{}
	if err := json.Unmarshal(userInfo, &socialUser); err != nil {
		return c.String(500, err.Error())
	}

	log.Info("소셜 로그인으로 user 정보 획득", string(userInfo))

	// google은 {"error":{""}} token invalid 하면 이렇게 응답 옴
	if socialUser["error"] != nil {
		return c.String(401, "invalid token")
	}

	if socialUser["id"] == nil {
		log.Error("socialUser 값이 존재 하지 않음...")
		return c.String(500, "알 수 없는 오류가 발생했습니다.")
	}

	// 소셜로그인으로 id 값 획득하고 소셜 id로 UserAuth 테이블 탐색 후 존재하지 않으면 해당 userId에 social_id 저장
	socialId := socialUser["id"].(string)
	userAuth, err := model.GetUserAuthBySocialId(socialId)
	if err != nil {
		// does not exist
		if err == sql.ErrNoRows {
			err := model.SetUserAuthWithSocialId(user["user_id"], socialId)
			if err != nil {
				return c.String(500, err.Error())
			}
		} else {
			return c.String(500, err.Error())
		}
	}

	// UserAuth의 user_id, social_id 값과 request body로 요청 오는 user_id 값이 다르면 잘못된 요청
	if userAuth != nil {
		if user["user_id"] != userAuth.UserId {
			return c.String(401, "이미 가입한 유저입니다.")
		}
	}

	// UserAuth의 userId 값으로 앱 내에서 사용할 토큰 생성
	token, err := util.CreateAccessToken(user["user_id"])
	if err != nil {
		return c.String(500, err.Error())
	}

	res := map[string]string{
		"access_token": token,
	}

	return c.JSON(200, res)
}
