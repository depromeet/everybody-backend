package controller

import (
	"strconv"

	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func SendFeedback(c echo.Context) error {
	req := c.Request() // 들어온 원본 request를 획득하여 그대로 사용

	// 인증토큰이 존재하면 획득 후 삭제
	token := req.Header.Get("Authorization")
	if token != "" {
		req.Header.Del("Authorization")
	}

	// 인증절차 수행 미인증이면 id=0
	id, _ := util.VerifyAccessToken(token)
	req.Header.Set("user", strconv.Itoa(id))

	log.Info("SendFeedback")

	code, header, body := callRestApi(c, c.Request(), "/feedbacks", "POST")
	if header != nil {
		for k, v := range header {
			c.Response().Header().Set(k, v.(string))
		}
	}

	return c.String(code, body)
}