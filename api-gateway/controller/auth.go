package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
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
		panic(err)
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

	// access to DB
	ua := model.GetUserAuth(reqUa.UserId)
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

func SignUp(c echo.Context) error {
	// read json body
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		panic(err)
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
	password := reqMap["password"].(string)
	if password == "" {
		log.Error("password invalid...")
		return c.String(http.StatusBadRequest, "password invalid")
	}

	// rest-api 호출 - 사용자에게 받은 body를 그대로 넘김
	reqMapByte, err := json.Marshal(reqMap)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	req, _ := http.NewRequest("", "", bytes.NewReader(reqMapByte)) // method and url will be set bottom
	req.Header.Set("Content-Type", "application/json")
	code, _, resBody := callRestApi(c, req, "/users", "POST")

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
		ua.Password = password
		model.SetUserAuth(ua)
		return c.JSON(http.StatusOK, ua) // 성공 시 user_id와 password만 리턴

	} else { // restapi 서버에서 응답이 200이 아닌 경우, DB 접근x
		log.Error("rest-api Request fail... code=" + strconv.Itoa(code) + " resBody=" + resBody)
		return c.String(code, resBody) // 실패 시 rest-api로부터 받은 실패코드와 메세지를 그대로 리턴해줌
	}
}
