// service_test.go 는 service 계층의 테스트를 위한
// 공통적인 작업을 정의합니다.
package service

import (
	"testing"

	"github.com/depromeet/everybody-backend/rest-api/mocks"
)

var (
	deviceRepo       *mocks.DeviceRepository
	notificationRepo *mocks.NotificationRepository
	userRepo         *mocks.UserRepository
	albumRepo        *mocks.AlbumRepositoryInterface
	pictureRepo      *mocks.PictureRepositoryInterface
	videoRepo        *mocks.VideoRepositoryInterface
	pushAdapter      *mocks.PushAdapter
	videoPort        *mocks.VideoPort
)

// initialize 는 서비스 계층 이외의 것들을 초기화합니다.
// 주로 repository 계층을 mocking 합니다.
func initialize(t *testing.T) {
	deviceRepo = new(mocks.DeviceRepository)
	notificationRepo = new(mocks.NotificationRepository)
	userRepo = new(mocks.UserRepository)
	albumRepo = new(mocks.AlbumRepositoryInterface)
	pictureRepo = new(mocks.PictureRepositoryInterface)
	videoRepo = new(mocks.VideoRepositoryInterface)
	pushAdapter = new(mocks.PushAdapter)
	videoPort = new(mocks.VideoPort)
}
