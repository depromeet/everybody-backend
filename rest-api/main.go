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
	userRepo repository.UserRepository

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
	userRepo = repository.NewUserRepository(dbClient)
	userService = service.NewUserService(userRepo)
	userService = service.NewUserService(userRepo)
	userHandler = handler.NewUserHandler(userService)
	server = http.NewServer(userHandler)
}
