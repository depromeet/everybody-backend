package service

import (
	"errors"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

var (
	ErrDuplicatedUserID = errors.New("고유한 유저 ID 생성에 실패했습니다.")
	signUpDefaultNickname = "끈육몬"
	signUpDefaultNotificationInterval = 3
)

type UserService interface {
	SignUp(body *dto.SignUpRequest) (*ent.User, error)
	GetUser(id string) (*ent.User, error)
}

func NewUserService(
	userRepo repository.UserRepository,
	notificationService NotificationService,
	deviceService DeviceService) UserService {
	return &userService{
		userRepo: userRepo,
		notificationService: notificationService,
		deviceService: deviceService,
	}
}

// UserService 의 구현체
type userService struct {
	userRepo repository.UserRepository
	notificationService NotificationService
	deviceService DeviceService
}

// SignUp 는 유저 생성 후 해당 유저의 Device를 등록합니다.
// 유저의 ID는 랜덤하고 고유한 UUID 입니다.
// 닉네임은 정의되지 않은 경우 "끈육몬"이 됨.
// TODO: 트랜잭션 롤백이 안됨. 유저를 만들고 다른 것들을 만들다가 종료되면..?
func (s *userService) SignUp(body *dto.SignUpRequest) (*ent.User, error) {
	maxTry := 10
	id := ""

	log.Infof("회원 가입을 위해 고유한 UUID 찾기")
	for i := 0; i < maxTry; i++{
		randomID := uuid.NewV4().String()
		log.Infof("%d 차 시도: %s", i+1, randomID)
		_, err := s.userRepo.FindById(randomID)
		if err != nil {
			notFoundErr := &ent.NotFoundError{}
			// NotFoundError 올바른 경우임.
			if errors.As(err, &notFoundErr) {
				log.Infof("고유한 ID: %s", randomID)
				id = randomID
				break
			} else {
				// notFoundError가 아닌 다른 에러
				return nil, err
			}
		}
	}

	if len(id) == 0 {
		log.Error("고유한 ID 찾기를 최대 재시도 했지만 찾지 못함.")
		// 이렇게 해도 에러 wrapping이 되나
		return nil, ErrDuplicatedUserID
	}

	if len(body.Nickname) == 0 {
		body.Nickname = signUpDefaultNickname
	}

	user, err := s.userRepo.Create(&ent.User{
		ID:          id,
		Nickname:    body.Nickname,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("유저를 생성했습니다. User(id=%s)", user.ID)

	if body.NotificationInterval == 0 {
		body.NotificationInterval = signUpDefaultNotificationInterval
	}

	notificationConfig, err := s.notificationService.Configure(id, &dto.ConfigureNotificationRequest{
		Interval: body.NotificationInterval, // 기본값
		IsActivated: true,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("알림 설정을 생성했습니다. NotificaitonConfig(id=%d)", notificationConfig.ID)

	device, err := s.deviceService.Register(id, &dto.RegisterDeviceRequest{
		DeviceToken: body.DeviceToken,
		PushToken: body.PushToken,
		DeviceOS: body.DeviceOS,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("디바이스 정보를 생성했습니다. Device(id=%d)", device.ID)

	return user, nil
}

// GetUser 는 유저 정보를 조회합니다.
func (s *userService) GetUser(id string) (*ent.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}
