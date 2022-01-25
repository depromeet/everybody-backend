package controller

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/depromeet/everybody-backend/api-gateway/model"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type OauthRequest struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
	Token    string `json:"token"`
}

func OauthLogin(c echo.Context) error {
	oauthRequest := &OauthRequest{}
	if err := c.Bind(oauthRequest); err != nil {
		log.Error("json body parse error...", "\nerr=", err)
		return c.String(400, "잘못된 요청입니다")
	}

	ua, err := model.GetUserAuth(oauthRequest.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("user not found...")
			return c.String(401, "user not found...")
		}
		log.Error(err)
		return c.String(500, "알 수 없는 오류가 발생했습니다")
	}

	if oauthRequest.Password != ua.Password {
		log.Info("password unmatched...")
		return c.String(401, "password unmatched...")
	}

	var socialId string
	switch oauthRequest.Kind {
	case "KAKAO":
		kakaoUserInfo, err := util.GetKakaoUserInfo(oauthRequest.Token)
		if err != nil {
			// 일단 임시로 에러 처리 문자열로 비교해서 하고 나중에 wraping해서든 수정하자ㅠ
			if strings.Contains(err.Error(), "invalid token") {
				return c.String(401, err.Error())
			}
			return c.String(500, err.Error())
		}
		socialId = strconv.Itoa(kakaoUserInfo.Id)

	case "GOOGLE":
		googleUserInfo, err := util.GetGoogleUserInfo(oauthRequest.Token)
		if err != nil {
			// 일단 임시로 에러 처리 문자열로 비교해서 하고 나중에 wraping해서든 수정하자ㅠ
			if strings.Contains(err.Error(), "invalid token") {
				return c.String(401, err.Error())
			}
			return c.String(500, err.Error())
		}
		socialId = googleUserInfo.Id

	case "APPLE":
		appleUserInfo, err := util.ValidateAppleTokenAndGetUserInfo(oauthRequest.Token)
		if err != nil {
			if strings.Contains(err.Error(), "invalid token") {
				return c.String(401, err.Error())
			}
			return c.String(500, err.Error())
		}

		socialId = appleUserInfo.Id

	default:
		log.Error("올바르지 않은 Oauth kind: ", oauthRequest.Kind)
		return c.JSON(400, "잘못된 요청입니다")
	}

	// 소셜 id로 UserAuth 테이블 탐색 후 존재하지 않으면 해당 userId에 social_id 저장
	userAuth, err := model.GetUserAuthBySocialId(socialId, oauthRequest.Kind)
	if err != nil {
		// does not exist
		if err == sql.ErrNoRows {
			err := model.SetUserAuthWithSocial(oauthRequest.UserId, socialId, oauthRequest.Kind)
			if err != nil {
				return c.String(500, "알 수 없는 오류가 발생했습니다")
			}
		} else {
			return c.String(500, "알 수 없는 오류가 발생했습니다")
		}
	}

	log.Info("소셜id로 유저 조회 완료")
	// userAuth nil 검사하는 이유는 최초로 소셜 로그인 하는 경우 userAuth 값이 없기 때문
	// UserAuth의 user_id, social_id 값과 request body로 요청 오는 user_id 값이 다르면 다른 앱에서 접속한 것
	// 기존의 user_id로 매핑해야됨
	userId := oauthRequest.UserId
	if userAuth != nil {
		if oauthRequest.UserId != userAuth.UserId {
			log.Info("소셜 인증을 요청한 앱 유저 id -> ", oauthRequest.UserId, "\n기존 앱 유저 id-> ", userAuth.UserId)
			userId = userAuth.UserId
		}
	}

	// 앱 내에서 사용할 토큰 생성
	token, err := util.CreateAccessToken(userId)
	if err != nil {
		return c.String(500, "알 수 없는 오류가 발생했습니다.")
	}

	res := map[string]string{
		"access_token": token,
	}
	log.Info("Login ok... userId=", userId)

	return c.JSON(200, res)
}
