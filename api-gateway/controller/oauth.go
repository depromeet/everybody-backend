package controller

import (
	"crypto/rsa"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/model"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	googeAuthUrlAPI       = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	kakaoAuthUrlAPI       = "https://kapi.kakao.com/v2/user/me"
	appleAuthUrlPublicKey = "https://appleid.apple.com/auth/keys"
)

type OauthRequest struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
	Kind     string `json:"kind"`
	Token    string `json:"token"`
}

type GoogleUserInfoResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Error struct {
		Code    *int   `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

type KakaoUserInfoResponse struct {
	Id int `json:"id"`
	// HasSignedUp bool      `json:"has_signed_up"`
	ConnectedAt time.Time `json:"connected_at"`
	KakaoAcount struct {
		Email string `json:"email"`
	} `json:"kakao_account"`
	Msg  string `json:"msg"`
	Code *int   `json:"code"`
}

type AppleUserInfoResponse struct {
	Id    string
	Email string
}

// apple의 public keys는 한개 혹은 여러가지
type ApplePublicKey struct {
	Keys []struct {
		Kty string `json:"kty"`
		Kid string `json:"kid"`
		Use string `json:"use"`
		Alg string `json:"alg"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

func getApplePublicKey(url string) (*ApplePublicKey, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, nil
	}
	defer resp.Body.Close()

	publickKey := &ApplePublicKey{}
	if err := json.Unmarshal(data, publickKey); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return publickKey, nil
}

func validateAppleTokenAndGetUserInfo(token string) (*AppleUserInfoResponse, error) {
	publicKey, err := getApplePublicKey(appleAuthUrlPublicKey)
	if err != nil {
		return nil, err
	}

	tokenSplit := strings.Split(token, ".")
	signingString := tokenSplit[0] + "." + tokenSplit[1]
	signature := tokenSplit[2]
	decodeTokenHeader, err := base64.RawURLEncoding.DecodeString(tokenSplit[0])
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	jwtHeader := &struct {
		Alg string
		Kid string
	}{}

	err = json.Unmarshal(decodeTokenHeader, jwtHeader)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	// apple public keys는 3개가 오는데 그 중에서 id token의 header와 일치하는 값을 찾는다
	for _, key := range publicKey.Keys {
		if key.Kid == jwtHeader.Kid && key.Alg == jwtHeader.Alg {
			exponent, err := base64.RawURLEncoding.DecodeString(key.E)
			if err != nil {
				return nil, err
			}

			modules, err := base64.RawURLEncoding.DecodeString(key.N)
			if err != nil {
				return nil, err
			}

			publicKey := &rsa.PublicKey{
				N: big.NewInt(0).SetBytes(modules),
				E: int(big.NewInt(0).SetBytes(exponent).Uint64()),
			}

			if err := jwt.SigningMethodRS256.Verify(signingString, signature, publicKey); err != nil {
				return nil, err
			}

			// 그냥 jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{}, error){return publicKey, nil})
			// 이걸로 해도 될 듯
			// get payload from jwt-payload
			payload, err := base64.RawURLEncoding.DecodeString(tokenSplit[1])
			if err != nil {
				return nil, err
			}

			claims := &jwt.MapClaims{}
			if err := json.Unmarshal(payload, claims); err != nil {
				return nil, err
			}

			log.Info("id token의 payload: ", string(payload))
			log.Info("토큰 Verifying...")

			if (*claims).VerifyAudience(config.Config.Oauth.Apple.AppId, true) &&
				(*claims).VerifyExpiresAt(time.Now().Unix(), true) &&
				(*claims).VerifyIssuer("https://appleid.apple.com", true) { // TODO: nonce도 체크해줘야될까...
				// subject type이 아마 string 이긴 할텐데 일단 타입 체킹...
				switch socialId := (*claims)["sub"].(type) {
				case float64:
					return &AppleUserInfoResponse{
						Id: strconv.Itoa(int(socialId)),
					}, nil
				case json.Number:
					return &AppleUserInfoResponse{
						Id: socialId.String(),
					}, nil
				case string:
					return &AppleUserInfoResponse{
						Id: socialId,
					}, nil
				}
			}

			return nil, errors.New("invalid token")
		}
	}

	log.Info("public key의 kid, alg와 id-token의 kid, alg와 일치하지 않습니다 ", jwtHeader.Kid, ",", jwtHeader.Alg)
	return nil, errors.New("invalid token")
}

// https://developers.kakao.com/docs/latest/ko/reference/rest-api-reference#response-code
func getKakaoUserInfo(accessToken string) (*KakaoUserInfoResponse, error) {
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

	log.Info("카카오 로그인으로 user 정보 획득", string(userInfo))

	kakaoReponse := &KakaoUserInfoResponse{}
	if err := json.Unmarshal(userInfo, kakaoReponse); err != nil {
		return nil, errors.WithStack(err)
	}

	if kakaoReponse.Code != nil {
		log.Error("code: ", *kakaoReponse.Code, " msg: ", kakaoReponse.Msg)
		if *kakaoReponse.Code == -401 {
			// wraping 하고 싶은데...
			return nil, errors.New("invalid token")
		}

		return nil, errors.New("알 수 없는 오류가 발생했습니다")
	}

	return kakaoReponse, nil
}

func getGoogleUserInfo(accessToken string) (*GoogleUserInfoResponse, error) {
	resp, err := http.Get(googeAuthUrlAPI + accessToken)
	// accessToken이 만료되거나 invalid 할 때 error catch가 안되는 거 같은데 확인 필요
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	userInfo, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	log.Info("구글 로그인으로 user 정보 획득", string(userInfo))

	googleUserInfoRes := &GoogleUserInfoResponse{}
	if err := json.Unmarshal(userInfo, googleUserInfoRes); err != nil {
		log.Error("json body parse error...", "\nerr=", err)
		return nil, err
	}

	if googleUserInfoRes.Error.Code != nil {
		log.Error("code: ", *googleUserInfoRes.Error.Code, " msg: ", googleUserInfoRes.Error.Message)
		if *googleUserInfoRes.Error.Code == 401 {
			return nil, errors.New("invalid token")
		}

		return nil, errors.New("알 수 없는 오류가 발생했습니다")
	}

	return googleUserInfoRes, nil
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
		kakaoUserInfo, err := getKakaoUserInfo(oauthRequest.Token)
		if err != nil {
			// 일단 임시로 에러 처리 문자열로 비교해서 하고 나중에 wraping해서든 수정하자ㅠ
			if strings.Contains(err.Error(), "invalid token") {
				return c.String(401, err.Error())
			}
			return c.String(500, err.Error())
		}
		socialId = strconv.Itoa(kakaoUserInfo.Id)

	case "GOOGLE":
		googleUserInfo, err := getGoogleUserInfo(oauthRequest.Token)
		if err != nil {
			// 일단 임시로 에러 처리 문자열로 비교해서 하고 나중에 wraping해서든 수정하자ㅠ
			if strings.Contains(err.Error(), "invalid token") {
				return c.String(401, err.Error())
			}
			return c.String(500, err.Error())
		}
		socialId = googleUserInfo.Id

	case "APPLE":
		appleUserInfo, err := validateAppleTokenAndGetUserInfo(oauthRequest.Token)
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
