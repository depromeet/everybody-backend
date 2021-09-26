package handler

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserHandler struct {
	userService service.UserService
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	body := new(dto.RegisterRequest)
	err := ctx.BodyParser(body)
	if err != nil {
		return err
	}
	user, err := h.userService.Register(body)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	user, err := h.userService.GetUser(ctx.Params("id"))
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}