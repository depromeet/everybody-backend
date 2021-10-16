package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/depromeet/everybody-backend/rest-api/config"
	_ "github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/infra/http"
	"github.com/depromeet/everybody-backend/rest-api/infra/http/handler"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var (
	notificationRepo repository.NotificationRepository
	deviceRepo       repository.DeviceRepository
	userRepo         repository.UserRepository
	albumRepo        repository.AlbumRepositoryInterface
	pictureRepo      repository.PictureRepositoryInterface

	notificationService service.NotificationService
	deviceService       service.DeviceService
	userService         service.UserService
	albumService        service.AlbumServiceInterface
	pictureService      service.PictureServiceInterface

	userHandler         *handler.UserHandler
	notificationHandler *handler.NotificationHandler
	albumHandler        *handler.AlbumHandler
	pictureHandler      *handler.PictureHandler

	server *fiber.App
)

func main() {
	log.SetReportCaller(true)

	initialize()
	if err := server.Listen(":8888"); err != nil {
		log.Error(err)
	}
}

func initialize() {
	dbClient := repository.Connect()
	// Aws session 연결하는 부분
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		// enable AWS_SDK_LOAD_CONFIG
		SharedConfigState: session.SharedConfigEnable,
		// Profile name
		Profile: config.Config.AWS.Profile,
	}))
	_, err := sess.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	notificationRepo = repository.NewNotificationRepository(dbClient)
	deviceRepo = repository.NewDeviceRepository(dbClient)
	userRepo = repository.NewUserRepository(dbClient)
	albumRepo = repository.NewAlbumRepository(dbClient)
	pictureRepo = repository.NewPictureRepository(dbClient)

	notificationService = service.NewNotificationService(notificationRepo)
	deviceService = service.NewDeviceService(deviceRepo)
	userService = service.NewUserService(userRepo, notificationService, deviceService)
	albumService = service.NewAlbumService(albumRepo, pictureRepo)
	pictureService = service.NewPictureService(pictureRepo, sess)

	userHandler = handler.NewUserHandler(userService)
	notificationHandler = handler.NewNotificationHandler(notificationService)
	albumHandler = handler.NewAlbumHandler(albumService)
	pictureHandler = handler.NewPictureHandler(pictureService)

	server = http.NewServer(userHandler, notificationHandler, albumHandler, pictureHandler)
}
