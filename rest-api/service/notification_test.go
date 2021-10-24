package service

import (
	"testing"
)

// device 관련 테스트를 하기 위한 초기화 작업을 수행합니다.
func initializeNotificationTest(t *testing.T) *notificationService {
	initialize(t)
	// 생성자를 이용하면서도 테스트 진행 시에는
	// 추상화된 interface가 아닌 concrete한 타입을 이용하기 위함.
	return NewNotificationService(notificationRepo).(*notificationService)
}

//func TestDeviceService_Register(t *testing.T) {
//	t.Run("이미 같은 device token 정보가 존재하는 경우 기존 정보를 그대로 리턴", func(t *testing.T) {
//		stubDevice := new(ent.Device)
//		deviceSvc := initializeDeviceTest(t)
//		deviceRepo.On("FindByDeviceToken", mock.AnythingOfType("string")).Return(stubDevice, nil).Once()
//
//		device, err := deviceSvc.Register(1, &dto.RegisterDeviceRequest{})
//		assert.NoError(t, err)
//		assert.Equal(t, stubDevice, device)
//	})
//
//	t.Run("새로운 device 정보 생성", func(t *testing.T) {
//		deviceSvc := initializeDeviceTest(t)
//		deviceRepo.On("FindByDeviceToken", mock.AnythingOfType("string")).Return(nil, fmt.Errorf("테스트 에러: %w", new(ent.NotFoundError))).Once()
//		var tmpDevice *ent.Device
//		deviceRepo.On("CreateDevice", mock.AnythingOfType("*ent.Device")).Run(func(args mock.Arguments) {
//			tmpDevice = args.Get(0).(*ent.Device)
//		}).Return(tmpDevice, nil).Once()
//
//		_, err := deviceSvc.Register(1, &dto.RegisterDeviceRequest{
//			DeviceOS: "ANDROID",
//		})
//		assert.NoError(t, err)
//	})
//}
