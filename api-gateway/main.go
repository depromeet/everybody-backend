// ref: https://github.com/pangpanglabs/echosample/blob/master/main.go
package main

import (
	"strconv"

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
	e.Use(middleware.Logger()) // TODO: 요거 알아보기... + 에러처리 방안 - https://echo.labstack.com/guide/error-handling/

	// index page & health check api
	e.GET(config.Config.ApiGw.HealthCheckPath, (&controller.IndexController{}).Index)

	// auth apis
	e.POST("/auth/login", func(c echo.Context) error {
		return controller.Login(c)
	})
	e.POST("/auth/signup", func(c echo.Context) error {
		return controller.SignUp(c)
	})

	// picures apis
	e.POST("/pictures", func(c echo.Context) error {
		return controller.UploadPicture(c)
	})

	// fowarding to rest-api server apis
	// 위에서 method & uri 조합으로 걸리지 않았으면 이 아래 match any로 걸림 - ref: https://echo.labstack.com/guide/routing/#path-matching-order
	controller.RestApiController{}.Init(e.Group("/*"))

	// run server...
	if err := e.Start(":" + strconv.Itoa(config.Config.ApiGw.Port)); err != nil {
		log.Fatal(err)
	}
}
