package service

import (
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeviceService_Register(t *testing.T) {
	t.Run("이미 같은 device token 정보가 존재하는 경우 기존 정보를 그대로 리턴", func(t *testing.T) {
		stubDevice := new(ent.Device)
		before := func() {
			initialize()
			deviceRepository.On("FindByDeviceToken", mock.AnythingOfType("string")).Return(stubDevice, nil).Once()
		}
		before()

		device, err := deviceSvc.Register("", &dto.RegisterDeviceRequest{})
		assert.NoError(t, err)
		assert.Equal(t, stubDevice, device)
	})

	t.Run("새로운 device 정보 생성", func(t *testing.T) {
		before := func() {
			initialize()
			deviceRepository.On("FindByDeviceToken", mock.AnythingOfType("string")).Return(nil, fmt.Errorf("테스트 에러: %w", new(ent.NotFoundError))).Once()
			var tmpDevice *ent.Device
			deviceRepository.On("CreateDevice", mock.AnythingOfType("*ent.Device")).Run(func(args mock.Arguments) {
				tmpDevice = args.Get(0).(*ent.Device)
			}).Return(tmpDevice, nil).Once()

		}
		before()

		_, err := deviceSvc.Register("", &dto.RegisterDeviceRequest{
			DeviceOS: "ANDROID",
		})
		assert.NoError(t, err)
	})
}

func TestDeviceService_GetDevice(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		before := func() {
			initialize()
			deviceRepository.On("FindById", mock.AnythingOfType("int")).Return(&ent.Device{}, nil).Once()
		}
		before()

		device, err := deviceSvc.GetDevice(1)
		assert.NoError(t, err)
		assert.NotNil(t, device)
	})

	t.Run("존재하지 않는 Device", func(t *testing.T) {
		before := func() {
			initialize()
			deviceRepository.On("FindById", mock.AnythingOfType("int")).Return(nil, new(ent.NotFoundError)).Once()
		}
		before()

		device, err := deviceSvc.GetDevice(1)
		errNotFound := new(ent.NotFoundError)
		// TODO: 아직 Error wrapping을 적용 안해서 아주 간단한 테스트 로직임.
		// 추후에 개선할 것.
		assert.ErrorAs(t, err, &errNotFound)
		assert.Nil(t, device)
	})
}
