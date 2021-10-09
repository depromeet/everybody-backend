package handler

import (
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
)

func NewNotificationHandler(notificationService service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

type NotificationHandler struct {
	notificationService service.NotificationService
}

//func (h *NotificationHandler) GetConfig(ctx *fiber.Ctx) error {
//	configID, err := strconv.Atoi(ctx.Params("id"))
//	if err != nil {
//		return err
//	}
//
//	config, err := h.notificationService.GetConfig(configID)
//	if err != nil {
//		return err
//	}
//
//	return ctx.JSON(config)
//}

// GetConfig 는 요청 유저의 알림 설정을 조회합니다.
func (h *NotificationHandler) GetConfig(ctx *fiber.Ctx) error {
	user, err := util.GetRequestUserID(ctx)
	if err != nil {
		return err
	}

	config, err := h.notificationService.GetConfigByUser(user)
	if err != nil {
		return err
	}

	return ctx.JSON(config)
}
