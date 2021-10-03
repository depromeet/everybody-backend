package service

import (
	"errors"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrDuplicatedDevice = errors.New("같은 디바이스에 대한 정보가 이미 존재합니다.")
)

type DeviceService interface {
	Register(requestUser string, body *dto.RegisterDeviceRequest) (*ent.Device, error)
	GetDevice(id int) (*ent.Device, error)
}

func NewDeviceService(deviceRepo repository.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepo: deviceRepo,
	}
}

// notificationService 는 NotificationService 의 구현체
type deviceService struct {
	deviceRepo repository.DeviceRepository
}

// Configure 는 requestUser의 알림 설정을 설정(?)합니다.
func (s *deviceService) Register(requestUser string, body *dto.RegisterDeviceRequest) (*ent.Device, error) {
	// 생성
	d, err := s.deviceRepo.FindByDeviceToken(body.DeviceToken)
	if err != nil {
		//
		notFoundErr := &ent.NotFoundError{}
		if errors.As(err, &notFoundErr) {
			// 해당 device token의 기기가 존재하지 않으면 생성
			os := device.DeviceOs(body.DeviceOS)
			if err = device.DeviceOsValidator(os); err != nil {
				return nil, err
			}

			return s.deviceRepo.CreateDevice(&ent.Device{
				DeviceToken: body.DeviceToken,
				PushToken:   body.PushToken,
				DeviceOs:    os,
				Edges:       ent.DeviceEdges{User: &ent.User{ID: requestUser}},
			})
		}

		return nil, err
	}

	log.Warningf("이미 같은 디바이스를 이용하는 정보가 존재합니다. Device(id=%d)", d.ID)
	return d, nil
}

// GetConfig 는 알림 설정 정보를 조회합니다.
func (s *deviceService) GetDevice(id int) (*ent.Device, error) {
	device, err := s.deviceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return device, err
}
