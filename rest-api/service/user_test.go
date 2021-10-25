package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// 유저 관련 테스트 위한 초기화 작업을 수행합니다.
func initializeUserTest(t *testing.T) *userService {
	initialize(t)

	deviceSvc := new(mocks.DeviceService)
	notificationSvc := new(mocks.NotificationService)

	// TODO: device도 DTO로 리턴하도록 정의
	deviceSvc.On("Register", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.RegisterDeviceRequest")).Return(new(ent.Device), nil)
	notificationSvc.On("Configure", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.ConfigureNotificationRequest")).Return(new(dto.NotificationConfigDto), nil)

	return NewUserService(userRepo, notificationSvc, deviceSvc).(*userService)
}

func TestUserService_SignUp(t *testing.T) {
	userSvc := initializeUserTest(t)
	deviceSvc := userSvc.deviceService.(*mocks.DeviceService)
	notificationSvc := userSvc.notificationService.(*mocks.NotificationService)
	userRepo.On("Create", mock.AnythingOfType("*ent.User")).Return(&ent.User{}, nil)

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

	t.Run("오류) 회원가입 시 알림 설정 없음", func(t *testing.T) {
		user, err := userSvc.SignUp(&dto.SignUpRequest{
			Device: &dto.RegisterDeviceRequest{},
		})
		assert.ErrorIs(t, err, ErrMissingNotificationConfig)
		assert.Nil(t, user)
	})

	t.Run("오류) 회원가입 시 디바이스 정보 없음", func(t *testing.T) {
		user, err := userSvc.SignUp(&dto.SignUpRequest{
			NotificationConfig: &dto.ConfigureNotificationRequest{},
		})
		assert.ErrorIs(t, err, ErrMissingDevice)
		assert.Nil(t, user)
	})
}
