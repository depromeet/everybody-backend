package main

import (
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
	deviceRepo repository.DeviceRepository
	userRepo repository.UserRepository

	notificationService service.NotificationService
	deviceService service.DeviceService
	userService service.UserService


	userHandler *handler.UserHandler

	server *fiber.App
)

func main() {
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

	notificationService = service.NewNotificationService(notificationRepo)
	deviceService = service.NewDeviceService(deviceRepo)
	userService = service.NewUserService(userRepo, notificationService, deviceService)

	userHandler = handler.NewUserHandler(userService)
	server = http.NewServer(userHandler)
}
