package service

import (
	"github.com/depromeet/everybody-backend/rest-api/adapter/push"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type NotificationService interface {
	Configure(requestUser int, body *dto.ConfigureNotificationRequest) (*dto.NotificationConfigDto, error)
	GetConfig(id int) (*dto.NotificationConfigDto, error)
	GetConfigByUser(user int) (*dto.NotificationConfigDto, error)
	NotifyPeriodicNoonBody(errChan chan<- error)
}

// 아직 알림은 persistent layer와 상관 없음
// 그래서 그냥 가볍게 여기서 정의
type Notification struct {
	recipient int // 알림을 받을 유저
	device    *ent.Device
	title     string
	content   string
}

func NewNotificationService(
	notificationRepo repository.NotificationRepository,
	pushAdapter push.PushAdapter) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		pushAdapter:      pushAdapter,
	}
}

// notificationService 는 NotificationService 의 구현체
type notificationService struct {
	notificationRepo repository.NotificationRepository
	pushAdapter      push.PushAdapter
}

// Configure 는 requestUser의 알림 설정을 설정(?)합니다.
func (s *notificationService) Configure(requestUser int, body *dto.ConfigureNotificationRequest) (*dto.NotificationConfigDto, error) {
	// 생성
	config, err := s.notificationRepo.FindByUser(requestUser)
	n := time.Now()
	log.Warning(n.Hour(), n.Minute())
	if err != nil {
		errNotFound := new(ent.NotFoundError)
		if errors.As(err, &errNotFound) {
			var lastNotifiedAt *time.Time
			now := time.Now()
			// 알림 선호 시간이 지금을 지났으면 오늘 알림을 보내지 않도록 lastNotifiedAt에 값을 넣어준다.
			if body.PreferredTimeHour <= now.Hour() && body.PreferredTimeMinute < now.Minute() {
				lastNotifiedAt = &now
			}

			// 이 유저의 push 설정이 없으면 만든다.
			result, err := s.notificationRepo.CreateNotificationConfig(&ent.NotificationConfig{
				Monday:              body.Monday, // 기본값
				Tuesday:             body.Tuesday,
				Wednesday:           body.Wednesday,
				Thursday:            body.Thursday,
				Friday:              body.Friday,
				Saturday:            body.Saturday,
				Sunday:              body.Sunday,
				PreferredTimeHour:   body.PreferredTimeHour,
				PreferredTimeMinute: body.PreferredTimeMinute,
				LastNotifiedAt:      lastNotifiedAt,
				IsActivated:         body.IsActivated,
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
	var lastNotifiedAt *time.Time
	now := time.Now()

	// 알림 선호 시간이 지금을 지났으면 오늘 알림을 보내지 않도록 lastNotifiedAt에 값을 넣어준다.
	if body.PreferredTimeHour <= now.Hour() && body.PreferredTimeMinute < now.Minute() {
		lastNotifiedAt = &now
	}
	result, err := s.notificationRepo.Update(config.ID, &ent.NotificationConfig{
		Monday:              body.Monday, // 기본값
		Tuesday:             body.Tuesday,
		Wednesday:           body.Wednesday,
		Thursday:            body.Thursday,
		Friday:              body.Friday,
		Saturday:            body.Saturday,
		Sunday:              body.Sunday,
		PreferredTimeHour:   body.PreferredTimeHour,
		PreferredTimeMinute: body.PreferredTimeMinute,
		LastNotifiedAt:      lastNotifiedAt, // 이건 변하지 않는다.
		IsActivated:         body.IsActivated,
		Edges:               config.Edges, // 이건 변하지 않는다.
	})

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

// NotifyPeriodicNoonBody 는 정기적인 눈바디 사진 찍기알림을 보낸다.
func (s *notificationService) NotifyPeriodicNoonBody(errChan chan<- error) {
	// NotifyPeriodicNoonBody 만이 이 channel의 producer이므로
	// 얘가 꼭 close 해줘야한다.
	defer close(errChan)

	notificationConfigs, err := s.notificationRepo.FindAll()
	if err != nil {
		errChan <- errors.WithStack(err)
		return
	}

	for _, nc := range notificationConfigs {
		if nc.IsActivated == true {
			if s.needsNotify(nc) {
				logger := log.WithField("user", nc.Edges.User.ID)
				logger.Info("[START] 정기 눈바디 알림을 보냅니다.")
				_, err := s.notificationRepo.UpdateLastNotifiedAt(nc.ID, time.Now())
				if err != nil {
					errChan <- err
				} else {
					if nc.Edges.User == nil {
						log.Errorf("%s의 유저가 존재하지 않습니다.", nc)
					} else {
						devices := nc.Edges.User.Edges.Devices
						for _, device := range devices {
							if err := s.pushAdapter.Send("눈바디 찍는 날이에요!", "우리 같이 꾸준히 눈바디를 기록해나가요!", device); err != nil {
								errChan <- errors.WithStack(err)
							}
							logger.Infof("Device(pushToken=%s)에게 정기 눈바디 알림을 보냈습니다", device.PushToken[:int(math.Min(float64(len(device.PushToken)), 10))])
						}
					}
				}

				logger.Info("[FINISH] 정기 눈바디 알림을 보냈습니다.")

			}
		}
	}
}

func (s *notificationService) needsNotify(nc *ent.NotificationConfig) bool {
	// TODO: 우선 항상 알림 보내고 있음!
	if !nc.IsActivated {
		return false
	}
	now := time.Now()
	// 한 번도 알림을 받은 적 없다.
	if nc.LastNotifiedAt == nil {
		// 지금이 preferred랑 비교
		return nc.PreferredTimeHour <= now.Hour() && nc.PreferredTimeMinute <= now.Minute()
	}

	// 이미 오늘 보냈으면
	if nc.LastNotifiedAt.Year() == now.Year() &&
		nc.LastNotifiedAt.Month() == now.Month() &&
		nc.LastNotifiedAt.Day() == now.Day() {
		return false
	}

	if
	// 오늘이 활성화되어있는가
	(nc.Monday && now.Weekday() == time.Monday ||
		nc.Tuesday && now.Weekday() == time.Tuesday ||
		nc.Wednesday && now.Weekday() == time.Wednesday ||
		nc.Thursday && now.Weekday() == time.Thursday ||
		nc.Friday && now.Weekday() == time.Friday ||
		nc.Saturday && now.Weekday() == time.Saturday ||
		nc.Sunday && now.Weekday() == time.Sunday) &&
		// 한 번도 보낸 적이 없는가
		(nc.LastNotifiedAt == nil ||
			// 오늘 알림을 보낸 적이 없는가
			(nc.LastNotifiedAt.Year() != now.Year() ||
				nc.LastNotifiedAt.Month() != now.Month() ||
				nc.LastNotifiedAt.Day() != now.Day())) {
		return nc.PreferredTimeHour <= now.Hour() && nc.PreferredTimeMinute <= now.Minute()
	}

	return false
}
