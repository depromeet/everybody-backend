package main

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/adapter/push"
	"github.com/depromeet/everybody-backend/rest-api/config"
	_ "github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/infra/http"
	"github.com/depromeet/everybody-backend/rest-api/infra/http/handler"
	"github.com/depromeet/everybody-backend/rest-api/infra/routine"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

var (
	pushAdapter push.PushAdapter

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

	notifyRoutine *routine.NotifyRoutine
)

func main() {
	log.SetReportCaller(true)

	initialize()
	// eb 배포 테스트 위해 잠시 5000번으로 변경
	if err := server.Listen(":5000"); err != nil {
		log.Error(err)
	}
}

func initialize() {
	dbClient := repository.Connect()
	pushAdapter = push.NewFirebasePushAdapter()

	notificationRepo = repository.NewNotificationRepository(dbClient)
	deviceRepo = repository.NewDeviceRepository(dbClient)
	userRepo = repository.NewUserRepository(dbClient)
	albumRepo = repository.NewAlbumRepository(dbClient)
	pictureRepo = repository.NewPictureRepository(dbClient)

	notificationService = service.NewNotificationService(notificationRepo, pushAdapter)
	deviceService = service.NewDeviceService(deviceRepo)
	userService = service.NewUserService(userRepo, notificationService, deviceService)
	albumService = service.NewAlbumService(albumRepo, pictureRepo)
	pictureService = service.NewPictureService(pictureRepo)

	userHandler = handler.NewUserHandler(userService)
	notificationHandler = handler.NewNotificationHandler(notificationService)
	albumHandler = handler.NewAlbumHandler(albumService)
	pictureHandler = handler.NewPictureHandler(pictureService)

	if config.Config.NotifyRoutine.Enabled {
		notifyRoutine = routine.NewNotifyRoutine(pushAdapter, notificationService)

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
	}

	server = http.NewServer(userHandler, notificationHandler, albumHandler, pictureHandler)
}
