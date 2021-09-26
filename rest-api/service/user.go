package service

import (
	"errors"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	Register(body *dto.RegisterRequest) (*ent.User, error)
	GetUser(id string) (*ent.User, error)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// UserService 의 구현체
type userService struct {
	userRepo repository.UserRepository
}

// Register 는 random한 ID를 만들어내는 회원가입을 수행합니다.
// 초기에는 device token을 일종의 패스워드 개념으로 쓸 수도 있을 것 같아요.
func (u *userService) Register(body *dto.RegisterRequest) (*ent.User, error) {
	randomID := uuid.NewV4().String()

	_, err := u.userRepo.FindById(randomID)
	if err != nil {
		notFoundErr := &ent.NotFoundError{}
		// notFoundErr는 올바른 경우임.
		if errors.As(err, &notFoundErr) {
			//pass
		} else {
			log.Error(err)
			return nil, err
		}
	}

	user, err := u.userRepo.Create(&ent.User{
		ID:          randomID,
		Nickname:    "익명의 끄뉵잉",
		DeviceToken: body.DeviceToken,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser 는 유저 정보를 조회합니다.
func (u *userService) GetUser(id string) (*ent.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return user, err
}
