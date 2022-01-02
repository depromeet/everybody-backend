package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/depromeet/everybody-backend/rest-api/adapter/noti"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type FeedbackService interface {
	SendFeedback(sender int, body *dto.SendFeedbackRequest) error
}
type feedbackService struct {
	notifierPort noti.NotifierPort
}

func NewFeedbackService(notifierPort noti.NotifierPort) FeedbackService {
	return &feedbackService{
		notifierPort: notifierPort,
	}
}

func (s *feedbackService) SendFeedback(sender int, body *dto.SendFeedbackRequest) error {
	log.Infof("%s 가 피드백을 보냈습니다. %+v", strconv.Itoa(sender), body)
	starRateString := "별점 생략"
	if body.StarRate != nil {
		if *body.StarRate < 0 || 5 < *body.StarRate {
			starRateString = fmt.Sprintf("잘못된 별점(:%d)", *body.StarRate)
		} else {
			starRateString = strings.Repeat("★", int(*body.StarRate)) +
				strings.Repeat("☆", 5-int(*body.StarRate))
		}
	}

	msg := &noti.Message{
		Sender:  strconv.Itoa(sender),
		Title:   body.Title,
		Content: starRateString + "\n" + body.Content,
	}

	if err := s.notifierPort.Send(strconv.Itoa(sender), msg); err != nil {
		return errors.Wrap(err, "피드백을 보내지 못했습니다.")
	}

	return nil
}
