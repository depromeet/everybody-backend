package main

import (
	"github.com/depromeet/everybody-backend/rest-api/adapter/push"
	_ "github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/infra/http"
	"github.com/depromeet/everybody-backend/rest-api/infra/http/handler"
	"github.com/depromeet/everybody-backend/rest-api/infra/routine"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"time"
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

	pushAdapter push.PushAdapter
	server      *fiber.App

	notifyRoutine *routine.NotifyRoutine
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

	notificationRepo = repository.NewNotificationRepository(dbClient)
	deviceRepo = repository.NewDeviceRepository(dbClient)
	userRepo = repository.NewUserRepository(dbClient)
	albumRepo = repository.NewAlbumRepository(dbClient)
	pictureRepo = repository.NewPictureRepository(dbClient)

	notificationService = service.NewNotificationService(notificationRepo)
	deviceService = service.NewDeviceService(deviceRepo)
	userService = service.NewUserService(userRepo, notificationService, deviceService)
	albumService = service.NewAlbumService(albumRepo, pictureRepo)
	pictureService = service.NewPictureService(pictureRepo)

	userHandler = handler.NewUserHandler(userService)
	notificationHandler = handler.NewNotificationHandler(notificationService)
	albumHandler = handler.NewAlbumHandler(albumService)
	pictureHandler = handler.NewPictureHandler(pictureService)

	pushAdapter = push.NewFirebasePushAdapter()

	notifyRoutine = routine.NewNotifyRoutine(pushAdapter)

	// 우리 서버의 서브 루틴으로 알림 로직을 실행
	go func() {
		for {
			if err := notifyRoutine.Run(); err != nil {
				log.Errorf("NotifyRoutine 실행 도중 에러 발생: %+v", err)
				log.Errorf("잠시 후 다시 notifyRoutine을 실행합니다.")
				time.Sleep(time.Second)

			} else {
				log.Info("NotifyRoutine을 성공적으로 마쳤습니다.")
				break
			}

		}
	}()

	server = http.NewServer(userHandler, notificationHandler, albumHandler, pictureHandler)
}
