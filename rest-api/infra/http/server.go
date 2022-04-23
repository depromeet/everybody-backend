package http

import (
	"github.com/depromeet/everybody-backend/rest-api/infra/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewServer(
	userHandler *handler.UserHandler,
	notificationHandler *handler.NotificationHandler,
	albumHandler *handler.AlbumHandler,
	pictureHandler *handler.PictureHandler,
	videoHandler *handler.VideoHandler,
	feedbackHandler *handler.FeedbackHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandle,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(defaultLog)
	app.Get("/", index)

	addUserHandlers(app, userHandler)
	addNotificationHandlers(app, notificationHandler)
	addAlbumHandlers(app, albumHandler)
	addPictureHandlers(app, pictureHandler)
	addVideoHandlers(app, videoHandler)
	addFeedbackHandlers(app, feedbackHandler)

	return app
}

func index(ctx *fiber.Ctx) error {
	return ctx.SendString("눈바디를 쉽게 확인하도록 돕는 서비스 에브리바디의 RESTful API Server")
}

func addUserHandlers(app *fiber.App, userHandler *handler.UserHandler) {
	group := app.Group("/users")
	group.Post("", userHandler.SignUp)
	group.Get("/me", userHandler.GetUser)
	group.Put("/me", userHandler.UpdateUser)
	group.Put("/me/profile-image", userHandler.UpdateProfileImage)
	group.Put("/me/download-completed", userHandler.NotifyDownloadImage)
}

func addNotificationHandlers(app *fiber.App, notificationHandler *handler.NotificationHandler) {
	// 본인에 대한 알림 설정 조회만 수행
	app.Get("/notification-configs/me", notificationHandler.GetConfig)
	app.Put("/notification-configs/me", notificationHandler.UpdateConfig)
}

func addAlbumHandlers(app *fiber.App, albumHandler *handler.AlbumHandler) {
	group := app.Group("/albums")
	group.Post("", albumHandler.CreateAlbum)
	group.Get("", albumHandler.GetAllAlbums)
	group.Get("/:album_id", albumHandler.GetAlbum)
	group.Put("/:album_id", albumHandler.UpdateAlbum)
	group.Delete("/:album_id", albumHandler.DeleteAlbum)
}

func addPictureHandlers(app *fiber.App, pictureHandler *handler.PictureHandler) {
	group := app.Group("/pictures")
	group.Post("", pictureHandler.SavePicture)
	// query string으로 구분(uploader=?&album=?&body_part=?)
	group.Get("", pictureHandler.GetAllPictures)
	group.Get("/:picture_id", pictureHandler.GetPicture)
	group.Delete("/:picture_id", pictureHandler.DeletePicture)
}

func addVideoHandlers(app *fiber.App, videoHandler *handler.VideoHandler) {
	group := app.Group("/videos")
	group.Post("", videoHandler.SaveVideo)
	group.Get("", videoHandler.GetAllVideos)
	group.Get("/:video_id", videoHandler.GetVideo)
}

func addFeedbackHandlers(app *fiber.App, feedbackHandler *handler.FeedbackHandler) {
	group := app.Group("/feedbacks")
	group.Post("", feedbackHandler.SendFeedback)
}
