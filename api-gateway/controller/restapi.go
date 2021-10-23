package controller

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type RestApiController struct {
}

func (c RestApiController) Init(g *echo.Group) {
	handlerGet := func(c echo.Context) error {
		return callRestApi(c, "GET")
	}

	handlerPost := func(c echo.Context) error {
		return callRestApi(c, "POST")
	}

	handlerPut := func(c echo.Context) error {
		return callRestApi(c, "PUT")
	}

	handlerDelete := func(c echo.Context) error {
		return callRestApi(c, "DELETE")
	}

	g.GET("/*", handlerGet)
	g.POST("/*", handlerPost)
	g.PUT("/*", handlerPut)
	g.DELETE("/*", handlerDelete)
}

func callRestApi(c echo.Context, method string) error {
	req := c.Request() // http.Request 타입의 원본 req 획득

	// 인증토큰 검증 및 user_id로 변환 처리  // TODO: 미들웨어?로 빼는 방안 고려 필요
	if config.Config.ApiGw.AuthEnable {
		token := req.Header.Get("Authorization")
		if token == "" {
			log.Error("Token not exist")
			return c.String(http.StatusUnauthorized, "Token not exist") // TODO: 에러 리턴 방식 수정 필요
		}
		req.Header.Del("Authorization")
		// TODO: jwt 검증/해독/에러처리 + DB 연결해서 user_id 가져오기
		userId := strconv.Itoa(123) // ### temp....
		req.Header.Add("user", userId)
		log.Info("Token Decoding Success -> user_id=", userId)
	}
	req.Header.Add("user", "2") // ### temp.... TODO: this line should be deleted

	// 수신한 원본req의 destnation 조작
	targetAddress := config.Config.TargetServer.RestApi.Address
	targetPath := req.URL.String()[1+len(config.Config.TargetServer.RestApi.Prefix):] // ex. /restapi/aa/?myqs=777&myqs2=111 ---> /aa/?myqs=777&myqs2=111
	newURL, err := url.Parse(targetAddress + targetPath)
	if err != nil {
		log.Error(err)
		panic(err) // TODO: panic이 아니고 그냥 return 해서 요청 거절하면될듯??
	}
	req.URL = newURL
	req.Method = method
	req.RequestURI = "" // need reset.. ref: https://stackoverflow.com/questions/19595860/http-request-requesturi-field-when-making-request-in-go
	log.Info("callRestApi -> "+method+": ", newURL.String())

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
