// ref: https://github.com/pangpanglabs/echosample/blob/master/main.go
package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/controller"
)

func main() {
	log.Info("SERVER START.....")

	// TODO: DB init함수 구현필요
	/*
		db, err := initDB(c.Database.Driver, c.Database.Connection)
		if err != nil {
			panic(err)
		}
		defer db.Close()
	*/

	e := echo.New()

	// register API GATEWAY's api...
	controller.RestApiController{}.Init(e.Group("/restapi"))
	//controller.AuthController{}.Init(e.Group("/auth"))

	// register server health check api
	e.GET(config.Config.ApiGw.HealthCheckPath, func(c echo.Context) error { // TODO: 연동테스트 완료되면 * 제거
		c.Response().Header().Set("Health-Checked-Time", time.Now().Format(time.RFC3339))
		return c.String(http.StatusOK, "OK")
	})

	// register app version check api
	// TODO: TBD

	// run server...
	if err := e.Start(":" + strconv.Itoa(config.Config.ApiGw.Port)); err != nil {
		log.Fatal(err)
	}
}

func initDB() {
	// TODO:
}
