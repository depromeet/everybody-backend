package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/model"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type AuthController struct {
}

func Login(c echo.Context) error {
	// read json body
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		panic(err) // TODO: log.panic으로 바꾸고.. middleware recover 도입.. https://echo.labstack.com/middleware/recover/
	}
	d := json.NewDecoder(strings.NewReader(string(data)))
	d.UseNumber()
	var reqMap map[string]interface{}
	if err := d.Decode(&reqMap); err != nil {
		log.Error("json body parse error... bodyStr=", string(data), "\nerr=", err)
		return c.String(http.StatusBadRequest, "json body parse error...")
	}
	log.Info("Login -> body=", reqMap)

	// make UserAuth Obj
	var reqUa model.UserAuth
	userId, err := strconv.Atoi(reqMap["user_id"].(json.Number).String())
	if err != nil {
		log.Error("json body parse error... bodyStr=", string(data), "\nerr=", err)
		return c.String(http.StatusBadRequest, "json body parse error...")
	}
	reqUa.UserId = userId
	reqUa.Password = reqMap["password"].(string)

	// access to
	// TODO: error 처리가 안되어 있는데 error 처리 해주어야 함
	ua, _ := model.GetUserAuth(reqUa.UserId)
	if ua.UserId < 1 {
		log.Error("user not found... req=", reqUa)
		return c.String(http.StatusBadRequest, "user not found...")
	}

	// password 일치 여부 검사
	if reqUa.Password != ua.Password {
		log.Info("password unmatched... req=", reqUa)
		return c.String(http.StatusBadRequest, "password unmatched...")
	}

	// JWT 발급
	token, err := util.CreateAccessToken(ua.UserId)
	if err != nil {
		log.Error("token creation fail...")
		return c.String(http.StatusInternalServerError, "token creation fail...")
	}

	res := map[string]interface{}{
		"access_token": token,
	}
	log.Info("Login ok... userId=", ua.UserId)
	return c.JSON(http.StatusOK, res)
}

const (
	googeAuthUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	kakaoAuthUrlAPI
	appleAuthUrlAPI
)

func getUserInfo(authURL, accessToken string) ([]byte, error) {
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

func GoogleLogin(c echo.Context) error {
	var user map[string]int
	if err := c.Bind(&user); err != nil {
		return errors.WithStack(err)
	}

	accessToken := c.Get("oauthtoken").(string)
	userInfo, err := getUserInfo(googeAuthUrlAPI, accessToken)
	if err != nil {
		log.Error(err.Error())
		return errors.WithStack(err)
	}

	var socialUser map[string]interface{}
	if err := json.Unmarshal(userInfo, &socialUser); err != nil {
		return errors.WithStack(err)
	}

	log.Info("소셜 로그인으로 user 정보 획득", string(userInfo))

	// TODO: 만약 userinfo의 id(소셜id)가 DB에 존재하지 않다면 new user -> save(db에 save)
	// 기존의 자동가입회원(?)의 정보는 어떻게 처리하는게 좋을까...

	// 소셜로그인으로 id 값 획득하고 소셜 id로 UserAuth 테이블 탐색 후 존재하지 않으면 해당 userId에 social_id 저장
	socialId := socialUser["id"].(string)
	userAuth, err := model.GetUserAuthBySocialId(socialId)
	if err != nil {
		// does not exist
		if err == sql.ErrNoRows {
			err := model.SetUserAuthWithSocialId(user["id"], socialId)
			if err != nil {
				return errors.WithStack(err)
			}
		} else {
			return errors.WithStack(err)
		}
	}

	// UserAuth의 userId 값으로 앱 내에서 사용할 토큰 생성
	token, err := util.CreateAccessToken(userAuth.UserId)
	if err != nil {
		return errors.WithStack(err)
	}

	res := map[string]string{
		"access_token": token,
	}

	return c.JSON(200, res)
}

func SignUp(c echo.Context) error {
	// read json body
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		panic(err) // TODO: log.panic으로 바꾸고.. middleware recover 도입.. https://echo.labstack.com/middleware/recover/
	}
	d := json.NewDecoder(strings.NewReader(string(data)))
	d.UseNumber()
	var reqMap map[string]interface{}
	if err := d.Decode(&reqMap); err != nil {
		log.Error("json body parse error... bodyStr=", string(data), "\nerr=", err)
		return c.String(http.StatusBadRequest, "json body parse error...")
	}
	log.Info("SignUp -> body=", reqMap)

	// check password validation
	password := reqMap["password"]
	if password == nil || reflect.TypeOf(password) != reflect.TypeOf("") || password == "" {
		log.Error("password invalid... password=", password)
		return c.String(http.StatusBadRequest, "password invalid")
	}

	// rest-api 호출 - 사용자에게 받은 body를 그대로 넘김
	reqMapByte, err := json.Marshal(reqMap)
	if err != nil {
		log.Error(err)
		panic(err) // TODO: log.panic으로 바꾸고.. middleware recover 도입.. https://echo.labstack.com/middleware/recover/
	}
	req, _ := http.NewRequest("", "", bytes.NewReader(reqMapByte)) // method and url will be set bottom
	req.Header.Set("Content-Type", "application/json")
	code, header, resBody := callRestApi(c, req, "/users", "POST")

	// 성공인 경우, UserAuth 테이블에 삽입하고, user_id/password만 리턴해줌
	if code == http.StatusOK {
		d := json.NewDecoder(strings.NewReader(resBody))
		d.UseNumber()
		resMap := make(map[string]interface{})
		if err := d.Decode(&resMap); err != nil {
			log.Error("resBody parse error... resBody=", resBody, "\nerr=", err)
			return c.String(http.StatusBadRequest, "resBody parse error...")
		}

		// UserAuth 테이블에 삽입
		var ua model.UserAuth
		userId, err := strconv.Atoi(resMap["id"].(json.Number).String())
		if err != nil {
			log.Error("resBody parse error... resBody=", resBody, "\nerr=", err)
			return c.String(http.StatusBadRequest, "resBody parse error...")
		}
		ua.UserId = userId
		ua.Password = password.(string)
		model.SetUserAuth(ua)

	} else { // restapi 서버에서 응답이 200이 아닌 경우, DB 접근x
		log.Error("rest-api Request fail... code=" + strconv.Itoa(code) + " resBody=" + resBody)
	}

	// rest에게 받은 응답을 그대로 전달
	if header != nil {
		for k, v := range header {
			c.Response().Header().Set(k, v.(string))
		}
	}
	return c.String(code, resBody)
}
