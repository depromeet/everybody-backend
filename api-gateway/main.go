package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/labstack/echo"

	log "github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	// server healthcheck api
	e.GET("/health*", func(c echo.Context) error { // TODO: 연동테스트 완료되면 * 제거
		c.Response().Header().Set("health_header", "okok")
		c.Response().Header().Set("health_header22", "okok22")
		return c.String(http.StatusOK, "OK")
	})

	// rest-api call
	e.GET("/restapi/:uri", RestApiHttpGet)
	// TODO: POST, DELETE, PUT

	// sign in
	//e.POST("/auth/signin", SignIn)

	// app version checking
	// TODO: TBD

	e.Logger.Fatal(e.Start(":8443")) // localhost:8443
}

// TODO: 요 함수를 Service 레벨로 옮기기
func RestApiHttpGet(c echo.Context) error {
	req := c.Request() // http.Request 타입의 원본 req 획득

	// 인증토큰 검증 및 user_id로 변환 처리
	// TODO: 미들웨어?로 빼는 방안 고려 필요
	token := req.Header.Get("Authorization")
	if token == "" {
		log.Error("Token not exist")
		//return c.String(http.StatusUnauthorized, "Token not exist") // TODO: 에러 리턴 방식 수정 필요
	}
	req.Header.Del("Authorization")
	// TODO: jwt 검증/해독/에러처리 + DB 연결해서 user_id 가져오기
	userId := strconv.Itoa(123) // temp..
	req.Header.Add("user_id", userId)
	log.Info("Token Decoding Success -> user_id=", userId)

	// 수신한 원본req의 destnation 조작
	baseURL := config.Config.ServerAddr.RestApi
	url, err := url.Parse(baseURL + "/" + c.Param("uri"))
	if err != nil {
		log.Error(err)
		panic(err)
	}
	req.URL = url
	req.Method = "GET"
	req.RequestURI = "" // need reset.. ref: https://stackoverflow.com/questions/19595860/http-request-requesturi-field-when-making-request-in-go
	log.Info("RestApiHttpGet -> ", url.String())

	// HttpClient 로 실제 앱서버 호출
	client := &http.Client{}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer resp.Body.Close()

	// 응답 받은 결과 복사 - Header
	for k := range resp.Header { // TODO: 반복문 대신 통짜로 복사할순 없을까??
		c.Response().Header().Set(k, resp.Header.Get(k))
	}

	// 응답 받은 결과 복사 - Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	str := string(data)

	log.Debug(str)
	return c.String(resp.StatusCode, str)
}
