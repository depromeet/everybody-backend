package http

import (
	"github.com/depromeet/everybody-backend/rest-api/infra/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer(
	userHandler *handler.UserHandler,
	notificationHandler *handler.NotificationHandler,
	albumHander *handler.AlbumHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandle,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(logger.New())
	app.Get("/", index)

	addUserHandlers(app, userHandler)
	addNotificationHandlers(app, notificationHandler)
	addAlbumHandlers(app, albumHander)

	return app
}

func index(ctx *fiber.Ctx) error {
	return ctx.SendString("눈바디를 쉽게 확인하도록 돕는 서비스 에브리바디의 RESTful API Server")
}

func addUserHandlers(app *fiber.App, userHandler *handler.UserHandler) {
	group := app.Group("/users")
	group.Get("/:id", userHandler.GetUser)
	group.Post("", userHandler.SignUp)
}

func addNotificationHandlers(app *fiber.App, notificationHandler *handler.NotificationHandler) {
	// 본인에 대한 알림 설정 조회만 수행
	app.Get("notification-configs/me", notificationHandler.GetConfig)
}

func addAlbumHandlers(app *fiber.App, albumHandler *handler.AlbumHandler) {
	app.Get("/album", albumHandler.CreateAlbum)
}
