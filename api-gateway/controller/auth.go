package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/model"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type AuthController struct {
}

func (c AuthController) Init(g *echo.Group) {
	g.POST("/login", c.Login)
	g.POST("/signup", c.SignUp)
}

func (AuthController) Login(c echo.Context) error {
	data, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	var reqUa model.UserAuth
	json.Unmarshal([]byte(string(data)), &reqUa)

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
	log.Info("login ok... userId=", strconv.FormatUint(ua.UserId, 10)) // to decimal string
	return c.JSON(http.StatusOK, res)
}

func (AuthController) SignUp(c echo.Context) error {
	// req body의 password 파싱 후 유효성 확인
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		log.Error("json parse error", err)
		return c.String(http.StatusBadRequest, "json parse error")
	}
	password := json_map["password"]
	if password == nil || reflect.TypeOf(password).Kind() != reflect.String || password == "" {
		log.Error("password invalid")
		return c.String(http.StatusBadRequest, "password invalid")
	}

	// rest-api 호출
	url := config.Config.TargetServer.RestApi.Address + "/users"
	jsonData, err := json.Marshal(json_map)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	log.Info("callRestApi -> POST: ", url)

	client := &http.Client{}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer resp.Body.Close()

	// rest-api 호출 결과 처리
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	// 성공인 경우, UserAuth 테이블에 user_id/password 로우 추가하고, 성공과 함께 user_id 리턴
	if resp.StatusCode == http.StatusOK {
		bodyMap := make(map[string]interface{})
		d := json.NewDecoder(bytes.NewBuffer([]byte(string(data))))
		d.UseNumber() // "id" 키를 기본 float64가 아닌 uint64로 받기 위해..
		if err := d.Decode(&bodyMap); err != nil {
			panic(err)
		}

		// UserAuth 테이블에 삽입
		var ua model.UserAuth
		userId, _ := strconv.ParseUint(string(bodyMap["id"].(json.Number)), 10, 64)
		ua.UserId = userId
		ua.Password = password.(string)
		model.SetUserAuth(ua)
		return c.JSON(http.StatusOK, ua)

	} else { // restapi 서버에서 응답이 200이 아닌 경우, DB 접근 없이 실패 리턴
		log.Info(strconv.Itoa(resp.StatusCode) + "//" + string(data))
		return c.String(http.StatusInternalServerError, "rest-api Request fail...")
	}
}
