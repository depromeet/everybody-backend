// ref: https://github.com/pangpanglabs/echosample/blob/master/main.go
package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/controller"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

func main() {
	log.Info("SERVER START.....")

	// DB Connectivity test
	db := util.CreateDBConn()
	db.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", (&controller.IndexController{}).Index)
	// server health check api
	e.GET(config.Config.ApiGw.HealthCheckPath, func(c echo.Context) error {
		c.Response().Header().Set("Health-Checked-Time", time.Now().Format(time.RFC3339))
		log.Info("health check OK...")
		return c.String(http.StatusOK, "OK")
	})

	// auth apis
	e.POST("/auth/login", func(c echo.Context) error {
		return controller.Login(c)
	})
	e.POST("/auth/signup", func(c echo.Context) error {
		return controller.SignUp(c)
	})

	// fowarding to rest-api server apis
	controller.RestApiController{}.Init(e.Group("/*"))
	// run server...
	if err := e.Start(":" + strconv.Itoa(config.Config.ApiGw.Port)); err != nil {
		log.Fatal(err)
	}
}
