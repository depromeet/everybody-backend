package dto

import (
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"time"
)

type SignUpRequest struct {
	Password             string `json:"password"`
	Nickname             string `json:"nickname"`
	NotificationInterval int    `json:"notification_interval"`
	DeviceToken          string `json:"device_token"`
	PushToken            string `json:"push_token"`
	DeviceOS             string `json:"device_os"`
	Type                 string `json:"type"`
	SNSAccessToken       string `json:"sns_access_token"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
}

type UserDto struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Height    *int      `json:"height"`
	Weight    *int      `json:"weight"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToDto(src *ent.User) *UserDto {
	return &UserDto{
		ID:        src.ID,
		Nickname:  src.Nickname,
		Height:    src.Height,
		Weight:    src.Weight,
		Kind:      src.Type.String(),
		CreatedAt: src.CreatedAt,
	}
}
