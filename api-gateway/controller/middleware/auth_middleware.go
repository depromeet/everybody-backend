package middleware

import (
	"strings"

	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// 소셜 로그인 할때 클라이언트로부터 access token을 받아서 next handler로 넘겨주는 middleware
func OauthTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := c.Request().Header["Authorization"]
		if len(headers) == 0 {
			log.Error("Oauth token이 존재하지 않습니다.")
			return c.JSON(403, "Invalid...")
		}

		bearerToken := strings.Split(headers[0], "Bearer ")
		if len(bearerToken) == 2 {
			c.Set("oauthtoken", strings.TrimSpace(bearerToken[1]))
		}

		next(c)
		return nil
	}
}

// TODO: 나중에 token 검증하는 부분 전부 다 middleware로 빼야함...
func AuthVerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headers := c.Request().Header["Authorization"]
		if len(headers) == 0 {
			log.Error("jwt token이 존재하지 않습니다.")
			return c.JSON(403, "Invalid...")
		}

		bearerToken := strings.Split(headers[0], "Bearer ")
		if len(bearerToken) == 2 {
			userId, err := util.VerifyAccessToken(headers[0])
			if err != nil {
				log.Error("토큰 인가 실패")
				return c.JSON(403, err.Error())
			}

			c.Set("user_id", userId)
		}

		next(c)
		return nil
	}
}
