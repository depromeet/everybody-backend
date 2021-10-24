package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
)

type NotificationService interface {
	Configure(requestUser int, body *dto.ConfigureNotificationRequest) (*dto.NotificationConfigDto, error)
	GetConfig(id int) (*dto.NotificationConfigDto, error)
	GetConfigByUser(user int) (*dto.NotificationConfigDto, error)
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
func (s *notificationService) Configure(requestUser int, body *dto.ConfigureNotificationRequest) (*dto.NotificationConfigDto, error) {
	// 생성
	config, err := s.notificationRepo.FindByUser(requestUser)
	if err != nil {
		errNotFound := new(ent.NotFoundError)
		if errors.As(err, &errNotFound) {
			result, err := s.notificationRepo.CreateNotificationConfig(&ent.NotificationConfig{
				Interval:    body.Interval, // 기본값
				IsActivated: body.IsActivated,
				Edges: ent.NotificationConfigEdges{
					User: &ent.User{ID: requestUser},
				},
			})
			if err != nil {
				return nil, errors.WithStack(err)
			}

			return dto.NotificationConfigToDto(result), nil
		} else {
			return nil, errors.WithStack(err)
		}
	}

	// 수정
	_, err = s.notificationRepo.UpdateInterval(config.ID, body.Interval)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result, err := s.notificationRepo.UpdateIsActivated(config.ID, body.IsActivated)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.NotificationConfigToDto(result), nil
}

// GetConfig 는 알림 설정의 ID로 알림 설정 정보를 조회합니다.
// 아직은 유저와 Config가 1대1이기에 굳이 Config ID를 이용할 필요 없습니다!
func (s *notificationService) GetConfig(id int) (*dto.NotificationConfigDto, error) {
	config, err := s.notificationRepo.FindById(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.NotificationConfigToDto(config), err
}

// GetConfigByUser 는 유저의 ID로 알림 설정 정보를 조회합니다.
func (s *notificationService) GetConfigByUser(user int) (*dto.NotificationConfigDto, error) {
	config, err := s.notificationRepo.FindByUser(user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.NotificationConfigToDto(config), nil
}
