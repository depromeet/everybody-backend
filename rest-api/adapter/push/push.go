package push

import "github.com/depromeet/everybody-backend/rest-api/ent"

type PushAdapter interface {
	// TODO: 여기는 device를 entity 형으로 받아야할까 dto형으로 받아야할까..?
	Send(title, content string, device *ent.Device) error
}
