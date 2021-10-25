package service

import (
	"testing"
)

// device 관련 테스트를 하기 위한 초기화 작업을 수행합니다.
func initializeNotificationTest(t *testing.T) *notificationService {
	initialize(t)
	// 생성자를 이용하면서도 테스트 진행 시에는
	// 추상화된 interface가 아닌 concrete한 타입을 이용하기 위함.
	return NewNotificationService(notificationRepo, pushAdapter).(*notificationService)
}

func TestNotificationService_NotifyPeriodicNoonBody(t *testing.T) {
	// TODO(umi0410): 주기적인 알림 테스트
}
