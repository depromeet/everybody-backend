package service

import (
	"errors"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrDuplicatedUserID               = errors.New("고유한 유저 ID 생성에 실패했습니다.")
	signUpDefaultNickname             = "끈육몬"
	signUpDefaultNotificationInterval = 3
)

type UserService interface {
	SignUp(body *dto.SignUpRequest) (*dto.UserDto, error)
	GetUser(id int) (*dto.UserDto, error)
}

func NewUserService(
	userRepo repository.UserRepository,
	notificationService NotificationService,
	deviceService DeviceService) UserService {
	return &userService{
		userRepo:            userRepo,
		notificationService: notificationService,
		deviceService:       deviceService,
	}
}

// UserService 의 구현체
type userService struct {
	userRepo            repository.UserRepository
	notificationService NotificationService
	deviceService       DeviceService
}

// SignUp 는 유저 생성 후 해당 유저의 Device를 등록합니다.
// 유저의 ID는 랜덤하고 고유한 UUID 입니다.
// 닉네임은 정의되지 않은 경우 "끈육몬"이 됨.
// TODO: 트랜잭션 롤백이 안됨. 유저를 만들고 다른 것들을 만들다가 종료되면..?
func (s *userService) SignUp(body *dto.SignUpRequest) (*dto.UserDto, error) {
	if len(body.Nickname) == 0 {
		body.Nickname = signUpDefaultNickname
	}

	user, err := s.userRepo.Create(&ent.User{
		Nickname: body.Nickname,
		Type:     user.Type(body.Type),
	})
	if err != nil {
		return nil, err
	}
	log.Infof("유저를 생성했습니다. User(id=%d)", user.ID)

	if body.NotificationInterval == 0 {
		body.NotificationInterval = signUpDefaultNotificationInterval
	}

	_, err = s.notificationService.Configure(user.ID, &dto.ConfigureNotificationRequest{
		Interval:    body.NotificationInterval, // 기본값
		IsActivated: true,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("알림 설정을 생성했습니다. NotificaitonConfig(user.id=%d)", user.ID)

	device, err := s.deviceService.Register(user.ID, &dto.RegisterDeviceRequest{
		DeviceToken: body.DeviceToken,
		PushToken:   body.PushToken,
		DeviceOS:    body.DeviceOS,
	})
	if err != nil {
		return nil, err
	}
	log.Infof("디바이스 정보를 생성했습니다. Device(id=%d)", device.ID)

	return dto.UserToDto(user), err
}

// GetUser 는 유저 정보를 조회합니다.
func (s *userService) GetUser(id int) (*dto.UserDto, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return dto.UserToDto(user), err
}
