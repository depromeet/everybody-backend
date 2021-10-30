package controller

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type RestApiController struct {
}

func (c RestApiController) Init(g *echo.Group) {
	// main에서 걸리지 않은 모든 path는 rest-api로 포워딩
	g.GET("/*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "GET")
	})
	g.POST("/*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "POST")
	})
	g.PUT("/*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "PUT")
	})
	g.DELETE("/*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "DELETE")
	})

}

func forwardWithAuthProc(c echo.Context, path string, method string) error {
	req := c.Request() // 들어온 원본 request를 획득하여 그대로 사용

	// 인증토큰이 존재하면 획득 후 삭제
	token := req.Header.Get("Authorization")
	if token != "" {
		req.Header.Del("Authorization")
	}

	// 인증절차 수행
	userId := uint64(0)
	if config.Config.ApiGw.AuthEnable {
		if token == "" {
			log.Error("Token not exist")
			return c.String(http.StatusUnauthorized, "Token not exist") // TODO: 에러 리턴 방식 수정 필요
		}
		id, err := util.VerifyAccessToken(token)
		if err != nil {
			return c.String(http.StatusForbidden, "Token invalid")
		}
		userId = id
		req.Header.Add("user", strconv.FormatUint(userId, 10)) // to decimal string
	} else {
		log.Warn("Auth DISABLED... func 'forwardToRestApi' processing with userId=0")
		req.Header.Add("user", "0") // 인증체크를 안하는 경우 userId를 0으로 넘김..
	}

	log.Info("forwardWithAuthProc -> userId=" + strconv.FormatUint(userId, 10) + " path=" + path)

	return c.String(callRestApi(c, false, c.Request(), path, method)) // TODO: true로 바꾸고 싶은데.. 바꾸면 헤더에 접근해버리고, 그러면 바디를 못넣어주는듯?
}

func callRestApi(c echo.Context, copyHeader bool, req *http.Request, path string, method string) (code int, s string) {
	// req의 destnation 설정
	newURL, err := url.Parse(config.Config.TargetServer.RestApi.Address + path) // path example: /aa/?myqs=777&myqs2=111
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, "url parse error"
	}
	req.URL = newURL
	req.Method = method
	req.RequestURI = "" // need reset.. ref: https://stackoverflow.com/questions/19595860/http-request-requesturi-field-when-making-request-in-go
	log.Info("callRestApi -> HTTP "+method+": ", newURL.String())

	// HttpClient 로 실제 앱서버 호출
	client := &http.Client{}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error(err)
		panic(err)
	}
	defer resp.Body.Close()

	// 응답 받은 결과 복사 - Header
	if copyHeader {
		for k := range resp.Header { // TODO: 반복문 대신 통짜로 복사할순 없을까??
			c.Response().Header().Set(k, resp.Header.Get(k))
		}
	}

	// 응답 받은 결과 복사 - Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	str := string(data)

	log.Debug(str)
	return resp.StatusCode, str
}
