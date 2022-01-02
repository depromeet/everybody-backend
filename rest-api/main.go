package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/adapter/noti"
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

func init() {
	rand.Seed(time.Now().Unix())
	// timezone을 global하게 잡아주지 않으면
	// local에선 KST를 사용하지만, EC2는 UTC를 사용한다.
	if loc, err := time.LoadLocation("Asia/Seoul"); err != nil {
		log.Fatal(err)
	} else {
		time.Local = loc
		log.Infof("Timezone을 설정합니다: %+v", time.Local)
	}
}

var (
	pushAdapter  push.PushAdapter
	notifierPort noti.NotifierPort

	notificationRepo repository.NotificationRepository
	deviceRepo       repository.DeviceRepository
	userRepo         repository.UserRepository
	albumRepo        repository.AlbumRepositoryInterface
	pictureRepo      repository.PictureRepositoryInterface
	videoRepo        repository.VideoRepositoryInterface

	notificationService service.NotificationService
	deviceService       service.DeviceService
	userService         service.UserService
	albumService        service.AlbumServiceInterface
	pictureService      service.PictureServiceInterface
	videoService        service.VideoServiceInterface
	feedbackService     service.FeedbackService

	userHandler         *handler.UserHandler
	notificationHandler *handler.NotificationHandler
	albumHandler        *handler.AlbumHandler
	pictureHandler      *handler.PictureHandler
	videoHandler        *handler.VideoHandler
	feedbackHandler     *handler.FeedbackHandler

	server *fiber.App

	notifyRoutine *routine.NotifyRoutine
)

func main() {
	initialize()
	if err := server.Listen(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
		log.Error(err)
	}
}

func initialize() {
	initializeLogger()

	dbClient := repository.Connect()
	pushAdapter = push.NewFirebasePushAdapter()
	notifierPort := noti.NewSlackNotifierAdapter()

	notificationRepo = repository.NewNotificationRepository(dbClient)
	deviceRepo = repository.NewDeviceRepository(dbClient)
	userRepo = repository.NewUserRepository(dbClient)
	albumRepo = repository.NewAlbumRepository(dbClient)
	pictureRepo = repository.NewPictureRepository(dbClient)
	videoRepo = repository.NewVideoRepository(dbClient)

	notificationService = service.NewNotificationService(notificationRepo, pushAdapter)
	deviceService = service.NewDeviceService(deviceRepo)
	userService = service.NewUserService(userRepo, notificationService, deviceService)
	albumService = service.NewAlbumService(albumRepo, pictureRepo)
	pictureService = service.NewPictureService(pictureRepo, albumRepo)
	feedbackService = service.NewFeedbackService(notifierPort)

	userHandler = handler.NewUserHandler(userService)
	notificationHandler = handler.NewNotificationHandler(notificationService)
	albumHandler = handler.NewAlbumHandler(albumService)
	pictureHandler = handler.NewPictureHandler(pictureService)
	videoHandler = handler.NewVideoHandler(videoService)
	feedbackHandler = handler.NewFeedbackHandler(feedbackService)

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

	server = http.NewServer(userHandler, notificationHandler, albumHandler, pictureHandler, videoHandler, feedbackHandler)
}
