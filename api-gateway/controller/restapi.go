package controller

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type RestApiController struct {
}

func (c RestApiController) Init(g *echo.Group) {
	// picures apis
	g.POST("pictures", func(c echo.Context) error {
		return UploadPicture(c)
	})
	g.GET("*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "GET")
	})
	g.POST("*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "POST")
	})
	g.PUT("*", func(c echo.Context) error {
		return forwardWithAuthProc(c, c.Request().URL.String(), "PUT")
	})
	g.DELETE("*", func(c echo.Context) error {
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
	userId := 0
	if config.Config.ApiGw.AuthEnable {
		if token == "" {
			log.Error("Token not exist...")
			return c.String(http.StatusUnauthorized, "Token not exist...")
		}
		id, err := util.VerifyAccessToken(token)
		if err != nil {
			log.Error("Token invalid... err=", err)
			return c.String(http.StatusForbidden, "Token invalid...")
		}
		userId = id

	} else {
		log.Warn("Auth DISABLED... func 'forwardToRestApi' processing with userId=0")
	}
	req.Header.Set("user", strconv.Itoa(userId))

	log.Info("forwardWithAuthProc -> userId=" + strconv.Itoa(userId) + " path=" + path)

	code, header, body := callRestApi(c, c.Request(), path, method)
	for k := range header {
		c.Response().Header().Set(k, header[k].(string))
	}
	return c.String(code, body)
}

func callRestApi(c echo.Context, req *http.Request, path string, method string) (int, map[string]interface{}, string) {
	// req의 destnation 설정
	newURL, err := url.Parse(config.Config.TargetServer.RestApi.Address + path) // path example: /aa/?myqs=777&myqs2=111
	if err != nil {
		log.Error(err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	req.URL = newURL
	req.Method = method
	req.RequestURI = "" // need reset.. ref: https://stackoverflow.com/questions/19595860/http-request-requesturi-field-when-making-request-in-go

	log.Info("callRestApi -> HTTP "+method+": ", newURL.String())

	// HttpClient 로 실제 앱서버 호출
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error("call rest-api fail... err=", err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	defer resp.Body.Close()

	// 응답 받은 결과 복사 - Header
	h := make(map[string]interface{})
	for k := range resp.Header {
		h[k] = resp.Header.Get(k)
	}

	// 응답 받은 결과 복사 - Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp.Body read error... err=", err)
		panic(err)
	}
	return resp.StatusCode, h, string(data)
}
