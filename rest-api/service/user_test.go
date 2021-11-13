package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// 유저 관련 테스트 위한 초기화 작업을 수행합니다.
func initializeUserTest(t *testing.T) *userService {
	initialize(t)

	deviceSvc := new(mocks.DeviceService)
	notificationSvc := new(mocks.NotificationService)

	return NewUserService(userRepo, notificationSvc, deviceSvc).(*userService)
}

func TestUserService_SignUp(t *testing.T) {
	userSvc := initializeUserTest(t)
	deviceSvc := userSvc.deviceService.(*mocks.DeviceService)
	notificationSvc := userSvc.notificationService.(*mocks.NotificationService)
	// TODO: device도 DTO로 리턴하도록 정의
	deviceSvc.On("Register", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.RegisterDeviceRequest")).Return(new(ent.Device), nil)
	notificationSvc.On("Configure", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.ConfigureNotificationRequest")).Return(new(dto.NotificationConfigDto), nil)

	userRepo.On("Create", mock.AnythingOfType("*ent.User")).Return(&ent.User{}, nil)
	userRepo.On("FindByNicknameContainingOrderByNicknameDesc", mock.AnythingOfType("string")).Return(nil, &ent.NotFoundError{})

	t.Run("성공) 회원가입", func(t *testing.T) {
		user, err := userSvc.SignUp(&dto.SignUpRequest{
			NotificationConfig: &dto.ConfigureNotificationRequest{},
			Device:             &dto.RegisterDeviceRequest{},
		})
		assert.NoError(t, err)

		// 회원가입 시에 수행되는 작업들을 확인한다.
		// 1. 유저 정보 생성
		userRepo.AssertNumberOfCalls(t, "Create", 1)
		// 2. 디바이스 등록
		deviceSvc.AssertNumberOfCalls(t, "Register", 1)
		// 3. 알림 설정 등록
		notificationSvc.AssertNumberOfCalls(t, "Configure", 1)
		t.Logf("유저 생성 %#v", user)
	})

	t.Run("성공) 회원가입 시 알림 설정 없으면 default config 적용", func(t *testing.T) {
		user, err := userSvc.SignUp(&dto.SignUpRequest{
			Device: &dto.RegisterDeviceRequest{},
		})
		assert.NoError(t, err)
		assert.NotNil(t, user)
	})

	t.Run("오류) 회원가입 시 디바이스 정보 없음", func(t *testing.T) {
		user, err := userSvc.SignUp(&dto.SignUpRequest{
			NotificationConfig: &dto.ConfigureNotificationRequest{},
		})
		assert.ErrorIs(t, err, ErrMissingDevice)
		assert.Nil(t, user)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("성공) 유저 정보 변경", func(t *testing.T) {
		userSvc := initializeUserTest(t)
		original := &ent.User{
			ID:        1,
			Nickname:  "beforeNickname",
			Motto:     "beforeMotto",
			Height:    nil,
			Weight:    nil,
			Kind:      "SIMPLE",
			CreatedAt: time.Now(),
		}
		var updated *ent.User = new(ent.User)

		userRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*ent.User")).Run(func(args mock.Arguments) {

			updateBody := args.Get(1).(*ent.User)
			// struct 복사. pointer X
			updated.ID = original.ID
			updated.CreatedAt = original.CreatedAt
			updated.Kind = original.Kind

			updated.Nickname = updateBody.Nickname
			updated.Motto = updateBody.Motto
			updated.Height = updateBody.Height
			updated.Weight = updateBody.Weight
		}).Return(updated, nil)
		height, weight := 170, 72
		result, err := userSvc.UpdateUser(original.ID, &dto.UpdateUserRequest{
			Nickname: "afterNickname",
			Motto:    "afterMotto",
			Height:   &height,
			Weight:   &weight,
		})
		assert.NoError(t, err)

		// 수정되어야하는 필드
		assert.Equal(t, result.ID, 1)
		assert.Equal(t, result.Nickname, "afterNickname")
		assert.Equal(t, result.Motto, "afterMotto")
		assert.Equal(t, *result.Height, 170)
		assert.Equal(t, *result.Weight, 72)

		// 수정되지 말아야하는 필드
		assert.Equal(t, result.CreatedAt, original.CreatedAt)
		assert.Equal(t, result.Kind, user.KindSIMPLE)
	})

	t.Run("실패) 존재하지 않는 유저에 대한 수정", func(t *testing.T) {
		// TODO(umi0410): repository에서 어떤 error를 발생하고 service가 어떻게 감쌀지는
		// 구체적으로 정해지지 않아서 대충 작성
		userSvc := initializeUserTest(t)
		userRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*ent.User")).
			Return(nil, errors.New("Entity not found"))
		_, err := userSvc.UpdateUser(999, &dto.UpdateUserRequest{})
		assert.NotNil(t, err)
	})
}
