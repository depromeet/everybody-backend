// TODO: 이 부분 패키지 명을 뭐로 하죠?
// 계속해서 loop를 돌면서 알림을 보내는 기능을 하는 녀석의 이름을 뭐로 붙이지
// 우선은 계속 이 루틴을 반복하니까 routine이라고 이름 붙여봤음.
// cron이라기에는 시각 기반 스케쥴링이 아니라 loop이고
// loop으로 이름 붙이자니 그냥 반복문 같은 느낌이라 좀 이상하고...
package routine

import (
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/adapter/push"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"time"
)

var (
	ErrNotifyRoutine = errors.New("NotifyRoutine에서 에러 발생")
)

// NotifyRoutine 는 유저에게 정기적인 알림을 보내는 루틴을 담당한다.
type NotifyRoutine struct {
	pushAdapter         push.PushAdapter
	notificationService service.NotificationService
}

func NewNotifyRoutine(adapter push.PushAdapter, notificationService service.NotificationService) *NotifyRoutine {
	return &NotifyRoutine{
		pushAdapter:         adapter,
		notificationService: notificationService,
	}
}

// Run 은 Routine이 계속해서 백그라운드에서 돌게 한다.
func (r *NotifyRoutine) Run() (err error) {
	for {
		log.Info("Notify Routine 루틴 시작")
		errChan := make(chan error)
		go func() {
			// panic 된 것들도 recover 해서 최대한 Run()이 종료되지 않게 함.
			// 각 goroutine에 대한 panic은 각 goroutine이 책임져야하나..?
			// => 헐 그런가보네... 처음 알았다;;;;;
			defer func() {
				if r := recover(); r != nil {
					log.Errorf("%+v:\n%s", r, debug.Stack())
					err = errors.Wrap(ErrNotifyRoutine, fmt.Sprintf("%s", r))
				} else {
					if err != nil {
						log.Errorf("%+v", err)
					}
				}
			}()
			r.notificationService.NotifyPeriodicNoonBody(errChan)
		}()
		for err := range errChan {
			log.Errorf("%+v", err)
		}
		log.Infof("Notify Routine Interval(%d sec) 동안 대기", config.Config.NotifyRoutine.Interval)
		// 개발용이라서 이렇게 짧게 쉬는데, 배포할 땐 1분 정도 잡아도 될 듯
		time.Sleep(time.Duration(config.Config.NotifyRoutine.Interval) * time.Second)
	}
}
