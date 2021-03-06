package push

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/device"
	"github.com/pkg/errors"
	"log"
	"os"
)

func init() {
	// GOOGLE_APPLICATION_CREDENTIALS 환경변수를 실제 환경변수로 이용하는 것은 번거로우니
	// config file의 path를 통해 설정.
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.Config.Push.FCM.ServiceAccountFile)
}

type firebasePushAdapter struct {
	firebase *firebase.App
}

// GOOGLE_APPLICATION_CREDENTIALS 환경 변수가 가리키는 location의 service account credential을 읽어와
// Firebase admin을 설정한다.
// 참고: https://github.com/firebase/firebase-admin-go/blob/master/snippets/init.go
func NewFirebasePushAdapter() PushAdapter {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &firebasePushAdapter{
		firebase: app,
	}
}

// 참고: https://github.com/firebase/firebase-admin-go/blob/master/snippets/messaging.go
func (a *firebasePushAdapter) Send(title, content string, deviceInfo *ent.Device) error {
	ctx := context.Background()
	client, err := a.firebase.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	registrationToken := deviceInfo.PushToken

	var response string
	fmt.Println(deviceInfo.DeviceToken)
	fmt.Println(device.DeviceOs(deviceInfo.DeviceToken))
	fmt.Println(device.DeviceOs(deviceInfo.DeviceToken) == device.DeviceOsANDROID)

	switch deviceInfo.DeviceOs {
	case device.DeviceOsANDROID:
		fallthrough
	case device.DeviceOsIOS:
		// See documentation on defining a message payload.
		message := newMessage(registrationToken, title, content)
		// Send a message to the device corresponding to the provided
		// registration token.
		response, err = client.Send(ctx, message)

	default:
		return errors.New("아직 지원되지 않는 OS. iOS와 ANDROID만 지원 중.")
	}

	// Response is a message ID string.
	fmt.Printf("푸시 알림 전송 완료: %s\n", response)

	return nil

}

func newMessage(token, title, content string) *messaging.Message {
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  content,
		},
		Token: token,
	}
	return message
}
