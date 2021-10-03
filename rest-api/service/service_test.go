package service

import (
	"github.com/depromeet/everybody-backend/rest-api/mocks"
)

var (
	deviceRepository *mocks.DeviceRepository
	deviceSvc        DeviceService
)

func initialize() {
	deviceRepository = new(mocks.DeviceRepository)
	deviceSvc = NewDeviceService(deviceRepository)
}
