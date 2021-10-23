package handler

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type UserHandler struct {
	userService service.UserService
}

func (h *UserHandler) SignUp(ctx *fiber.Ctx) error {
	body := new(dto.SignUpRequest)
	err := ctx.BodyParser(body)
	if err != nil {
		return errors.WithStack(err)
	}
	user, err := h.userService.SignUp(body)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(user)
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	id, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(user)
}
