package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
)

type NotificationService interface {
	Configure(requestUser string, body *dto.ConfigureNotificationRequest) (*ent.NotificationConfig, error)
	GetConfig(id int) (*ent.NotificationConfig, error)
}

func NewNotificationService(notificationRepo repository.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

// notificationService 는 NotificationService 의 구현체
type notificationService struct {
	notificationRepo repository.NotificationRepository
}

// Configure 는 requestUser의 알림 설정을 설정(?)합니다.
func (s *notificationService) Configure(requestUser string, body *dto.ConfigureNotificationRequest) (*ent.NotificationConfig, error) {
	// 생성
	if body.ID == 0 {
		return s.notificationRepo.CreateNotificationConfig(&ent.NotificationConfig{
			Interval: body.Interval, // 기본값
			IsActivated: body.IsActivated,
			Edges: ent.NotificationConfigEdges{
				User: &ent.User{ID: requestUser},
			},
		})
	}

	// 수정
	_, err := s.notificationRepo.UpdateInterval(body.ID, body.Interval)
	if err != nil {
		return nil, err
	}

	result, err := s.notificationRepo.UpdateIsActivated(body.ID, body.IsActivated)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetConfig 는 알림 설정 정보를 조회합니다.
func (s *notificationService) GetConfig(id int) (*ent.NotificationConfig, error){
	config, err := s.notificationRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return config, err
}
