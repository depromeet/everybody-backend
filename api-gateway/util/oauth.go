package util

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	googeAuthUrlAPI      = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	kakaoAuthUrlAPI      = "https://kapi.kakao.com/v2/user/me"
	applePublicKeyUrlAPI = "https://appleid.apple.com/auth/keys"
)

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

func ValidateAppleTokenAndGetUserInfo(IdToken string) (*AppleUserInfoResponse, error) {
	publicKey, err := getApplePublicKey(applePublicKeyUrlAPI)
	if err != nil {
		return nil, err
	}

	tokenSplit := strings.Split(IdToken, ".")
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

			log.Info("애플 로그인으로 user 정보 획득: ", string(payload))
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

	log.Error("public key의 kid, alg와 id-token의 kid, alg와 일치하지 않습니다 ", jwtHeader.Kid, ",", jwtHeader.Alg)
	return nil, errors.New("invalid token")
}

// https://developers.kakao.com/docs/latest/ko/reference/rest-api-reference#response-code
func GetKakaoUserInfo(accessToken string) (*KakaoUserInfoResponse, error) {
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

	log.Info("카카오 로그인으로 user 정보 획득")

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

func GetGoogleUserInfo(accessToken string) (*GoogleUserInfoResponse, error) {
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

	log.Info("구글 로그인으로 user 정보 획득")

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
