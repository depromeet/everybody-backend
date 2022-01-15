// ref: https://github.com/pangpanglabs/echosample/blob/master/main.go
package main

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/controller"
	am "github.com/depromeet/everybody-backend/api-gateway/controller/middleware"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

func init() {
	// timezone을 global하게 잡아주지 않으면
	// local에선 KST를 사용하지만, EC2는 UTC를 사용한다.
	if loc, err := time.LoadLocation("Asia/Seoul"); err != nil {
		log.Fatal(err)
	} else {
		time.Local = loc
		log.Infof("Timezone을 설정합니다: %+v", time.Local)
	}
}
func main() {
	log.Info("SERVER START.....")

	// DB Connectivity test
	db := util.CreateDBConn()
	db.Close()

	e := echo.New()
	e.Use(middleware.Logger()) // TODO: 요거 알아보기... + 에러처리 방안 - https://echo.labstack.com/guide/error-handling/
	e.Use(middleware.Recover())
	// index page & health check api
	e.GET(config.Config.ApiGw.HealthCheckPath, (&controller.IndexController{}).Index)

	// auth apis
	e.POST("/auth/login", func(c echo.Context) error {
		return controller.Login(c)
	})
	e.POST("/auth/signup", func(c echo.Context) error {
		return controller.SignUp(c)
	})

	// oauth := e.Group("/oauth", am.OauthTokenMiddleware)
	// oauth.POST("/google", controller.GoogleLogin)
	// oauth.POST("/kakao", controller.KakaoLogin)
	e.POST("/oauth/login", am.OauthTokenMiddleware(controller.OauthLogin))

	// picures apis
	e.POST("/pictures", func(c echo.Context) error {
		return controller.UploadPicture(c)
	})
	e.POST("/videos/download", func(c echo.Context) error {
		return controller.DownloadVideo(c)
	})

	// fowarding to rest-api server apis
	// 위에서 method & uri 조합으로 걸리지 않았으면 이 아래 match any로 걸림 - ref: https://echo.labstack.com/guide/routing/#path-matching-order
	controller.RestApiController{}.Init(e.Group("/*"))

	// run server...
	if err := e.Start(":" + strconv.Itoa(config.Config.ApiGw.Port)); err != nil {
		log.Fatal(err)
	}
}
