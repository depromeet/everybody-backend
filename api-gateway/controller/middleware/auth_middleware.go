package middleware

import (
	"net/http"
	"strings"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/util"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// 소셜 로그인 할때 클라이언트로부터 access token을 받아서 next handler로 넘겨주는 middleware
func OauthTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if len(header) == 0 {
			log.Error("Oauth token이 존재하지 않습니다.")
			return c.JSON(403, "Invalid...")
		}

		bearerToken := strings.Split(header, "Bearer ")
		if len(bearerToken) == 2 {
			c.Set("oauthtoken", strings.TrimSpace(bearerToken[1]))
		}

		next(c)
		return nil
	}
}

// TODO: 나중에 token 검증하는 부분 전부 다 middleware로 빼야함...
func AuthVerifyTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token != "" {
			c.Request().Header.Del("Authorization")
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
		c.Set("user", userId)
		next(c)
		return nil
	}
}
