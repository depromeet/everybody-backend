package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	ErrUnsupportedDevice = errors.New("지원되지 않는 기기 운영체제 입니다.")
)

type DeviceService interface {
	Register(requestUser int, body *dto.RegisterDeviceRequest) (*ent.Device, error)
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
func (s *deviceService) Register(requestUser int, body *dto.RegisterDeviceRequest) (*ent.Device, error) {
	// 동일한 디바이스 토큰 정보가 이미 존재하는지
	d, err := s.deviceRepo.FindByDeviceToken(body.DeviceToken)
	if err != nil {
		notFoundErr := new(ent.NotFoundError)
		if errors.As(err, &notFoundErr) {
			// 해당 device token의 기기가 존재하지 않으면 생성
			os := device.DeviceOs(body.DeviceOS)
			if err = device.DeviceOsValidator(os); err != nil {
				// validation error의 경우 특정 에러 타입이나 미리 정의된 에러 변수를 이용하지 않음.
				// ent 측에서 그냥 fmt.Errorf()로 에러를 생성해버림
				// 따라서 우리가 알아서 에러를 정의한 뒤 사용해야함
				return nil, errors.WithStack(ErrUnsupportedDevice)
			}
			// 생성
			return s.deviceRepo.CreateDevice(&ent.Device{
				DeviceToken: body.DeviceToken,
				PushToken:   body.PushToken,
				DeviceOs:    os,
				Edges:       ent.DeviceEdges{User: &ent.User{ID: requestUser}},
			})
		}

		return nil, errors.WithStack(err)
	}

	log.Warningf("이미 같은 디바이스를 이용하는 정보가 존재합니다. Device(id=%d)의 유저를 새 유저로 변경하겠습니다.", d.ID)
	update := d
	update.PushToken = body.PushToken
	update.Edges.User = &ent.User{ID: requestUser}
	err = s.deviceRepo.Update(d.ID, update)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return d, nil
}

// GetDevice 는 알림 설정 정보를 조회합니다.
func (s *deviceService) GetDevice(id int) (*ent.Device, error) {
	device, err := s.deviceRepo.FindById(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return device, nil
}
