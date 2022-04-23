package handler

import (
	"context"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func NewUserHandler(userService service.UserService, entClient *ent.Client) *UserHandler {
	return &UserHandler{
		userService: userService,
		entClient:   entClient,
	}
}

type UserHandler struct {
	userService service.UserService
	entClient   *ent.Client
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

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	id, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	body := new(dto.UpdateUserRequest)

	if err = ctx.BodyParser(body); err != nil {
		return errors.WithStack(err)
	}

	user, err := h.userService.UpdateUser(id, body)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(user)
}

func (h *UserHandler) UpdateProfileImage(ctx *fiber.Ctx) error {
	id, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	body := new(dto.UpdateProfileImageRequest)

	if err = ctx.BodyParser(body); err != nil {
		return errors.WithStack(err)
	}

	user, err := h.userService.UpdateProfileImage(id, body)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(user)
}

func (h *UserHandler) NotifyDownloadImage(ctx *fiber.Ctx) error {
	id, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	now := time.Now()
	update := h.entClient.User.UpdateOneID(id).SetNillableDownloadCompleted(&now)

	user, err := update.Save(context.Background())
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(user)
}
