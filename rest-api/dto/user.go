package dto

import (
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"time"
)

type SignUpRequest struct {
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Motto    string `json:"motto"`
	// 알림 설정할 때와 똑같은 body를 이용
	NotificationConfig *ConfigureNotificationRequest `json:"notification_config"`
	Device             *RegisterDeviceRequest        `json:"device"`
	// 유저 타입이 무엇인지 (type은 예약어이므로 kind를 사용)
	// e.g. SIMPLE, KAKAO, APPLE
	Kind           string `json:"kind"`
	Height         *int   `json:"height"`
	Weight         *int   `json:"weight"`
	SNSAccessToken string `json:"sns_access_token"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
	Motto    string `json:"motto"`
	Height   *int   `json:"height"`
	Weight   *int   `json:"weight"`
}

type UserDto struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Motto     string    `json:"motto"`
	Height    *int      `json:"height"`
	Weight    *int      `json:"weight"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
}

func UserToDto(src *ent.User) *UserDto {
	return &UserDto{
		ID:        src.ID,
		Motto:     src.Motto,
		Nickname:  src.Nickname,
		Height:    src.Height,
		Weight:    src.Weight,
		Kind:      src.Kind.String(),
		CreatedAt: src.CreatedAt,
	}
}
