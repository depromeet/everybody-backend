package handler

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func NewFeedbackHandler(feedbackService service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: feedbackService,
	}
}

type FeedbackHandler struct {
	feedbackService service.FeedbackService
}

// GetConfig 는 요청 유저의 알림 설정을 수정합니다.
func (h *FeedbackHandler) SendFeedback(ctx *fiber.Ctx) error {
	user, err := util.GetRequestUserID(ctx)
	log.Infof("미인증 유저의 피드백입니다. user을 0으로 설정합니다.")

	body := new(dto.SendFeedbackRequest)
	err = ctx.BodyParser(body)
	if err != nil {
		return errors.WithStack(err)
	}

	err = h.feedbackService.SendFeedback(user, body)
	if err != nil {
		log.Errorf("%+v", err)
		return ctx.Status(400).JSON(map[string]string{
			"message": "피드백을 전달하지 못했습니다.ㅜ^ㅜ\n잠시 후 다시 시도해주세요.",
		})
	}

	return ctx.JSON(map[string]string{
		"message": "피드백을 전달했습니다.",
	})
}
