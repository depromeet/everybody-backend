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

	deviceSvc.On("Register", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.RegisterDeviceRequest")).Return(new(ent.Device), nil)
	notificationSvc.On("Configure", mock.AnythingOfType("int"), mock.AnythingOfType("*dto.ConfigureNotificationRequest")).Return(new(ent.NotificationConfig), nil)

	return NewUserService(userRepo, notificationSvc, deviceSvc).(*userService)
}

func TestUserService_SignUp(t *testing.T) {
	userSvc := initializeUserTest(t)
	deviceSvc := userSvc.deviceService.(*mocks.DeviceService)
	notificationSvc := userSvc.notificationService.(*mocks.NotificationService)
	userRepo.On("Create", mock.AnythingOfType("*ent.User")).Return(&ent.User{}, nil)

	user, err := userSvc.SignUp(&dto.SignUpRequest{})
	assert.NoError(t, err)

	// 회원가입 시에 수행되는 작업들을 확인한다.
	// 1. 유저 정보 생성
	userRepo.AssertNumberOfCalls(t, "Create", 1)
	// 2. 디바이스 등록
	deviceSvc.AssertNumberOfCalls(t, "Register", 1)
	// 3. 알림 설정 등록
	notificationSvc.AssertNumberOfCalls(t, "Configure", 1)
	t.Logf("유저 생성 %#v", user)
}
