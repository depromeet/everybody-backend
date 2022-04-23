package dto

import (
	"fmt"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/ent/user"
)

type SignUpRequest struct {
	Password string `json:"password"`
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

type UpdateProfileImageRequest struct {
	ProfileImage string `json:"profile_image"`
}

type UserDto struct {
	ID                int        `json:"id"`
	Nickname          string     `json:"nickname"`
	Motto             string     `json:"motto"`
	Height            *int       `json:"height"`
	Weight            *int       `json:"weight"`
	Kind              user.Kind  `json:"kind"`
	ProfileImage      string     `json:"profile_image"`
	CreatedAt         time.Time  `json:"created_at"`
	DownloadCompleted *time.Time `json:"download_completed"`
}

func UserToDto(src *ent.User) *UserDto {

	return &UserDto{
		ID:                src.ID,
		Motto:             src.Motto,
		Nickname:          src.Nickname,
		Height:            src.Height,
		Weight:            src.Weight,
		Kind:              src.Kind,
		ProfileImage:      fmt.Sprintf("%s/%s", config.Config.PublicDriveRootURL, src.ProfileImage),
		CreatedAt:         src.CreatedAt,
		DownloadCompleted: src.DownloadCompleted,
	}
}
