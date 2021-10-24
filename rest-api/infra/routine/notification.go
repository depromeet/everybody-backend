// TODO: 이 부분 패키지 명을 뭐로 하죠?
// 계속해서 loop를 돌면서 알림을 보내는 기능을 하는 녀석의 이름을 뭐로 붙이지
// 우선은 계속 이 루틴을 반복하니까 routine이라고 이름 붙여봤음.
// cron이라기에는 시각 기반 스케쥴링이 아니라 loop이고
// loop으로 이름 붙이자니 그냥 반복문 같은 느낌이라 좀 이상하고...
package routine

import (
	"github.com/depromeet/everybody-backend/rest-api/adapter/push"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	ErrInternal = errors.New("NotificationRoutine에 알 수 없는 에러 발생")
)

// NotifyRoutine 는 유저에게 정기적인 알림을 보내는 루틴을 담당한다.
type NotifyRoutine struct {
	pushAdapter push.PushAdapter
}

func NewNotifyRoutine(adapter push.PushAdapter) *NotifyRoutine {
	return &NotifyRoutine{
		pushAdapter: adapter,
	}
}

// Run 은 Routine이 계속해서 백그라운드에서 돌게 한다.
func (r *NotifyRoutine) Run() (err error) {
	defer func() {
		// panic 된 것들도 recover 해서 최대한 Run()이 종료되지 않게 함.
		if r := recover(); r != nil {
			log.Error(r)
			err = errors.WithMessagef(ErrInternal, "%s", r)
		}
	}()

	for {
		log.Info("Running...")
		// 개발용이라서 이렇게 짧게 쉬는데, 배포할 땐 1분 정도 잡아도 될 듯
		time.Sleep(time.Second * 10)
	}
}
