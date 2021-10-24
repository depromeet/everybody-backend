package dto

import "github.com/depromeet/everybody-backend/rest-api/ent"

// clinet가 알림 설정을 설정할 때 보내는 Request
// 아직까지는 거의 NotificationConfigDto 와 동일
type ConfigureNotificationRequest struct {
	Monday              bool `json:"monday"`
	Tuesday             bool `json:"tuesday"`
	Wednesday           bool `json:"wednesday"`
	Thursday            bool `json:"thursday"`
	Friday              bool `json:"friday"`
	Saturday            bool `json:"saturday"`
	Sunday              bool `json:"sunday"`
	PreferredTimeHour   int  `json:"preferred_time_hour"`
	PreferredTimeMinute int  `json:"preferred_time_minute"`
	IsActivated         bool `json:"is_activated"`
}

// 알림 설정 DTO
type NotificationConfigDto struct {
	Monday              bool `json:"monday"`
	Tuesday             bool `json:"tuesday"`
	Wednesday           bool `json:"wednesday"`
	Thursday            bool `json:"thursday"`
	Friday              bool `json:"friday"`
	Saturday            bool `json:"saturday"`
	Sunday              bool `json:"sunday"`
	PreferredTimeHour   int  `json:"preferred_time_hour"`
	PreferredTimeMinute int  `json:"preferred_time_minute"`
	IsActivated         bool `json:"is_activated"`
}

// 진짜 필드 똑같음.. ㅎㅎ...
// 추후 확장할 때를 고려 + Entity를 직접 이용하기 보단 DTO를 이용하는 패턴을
// 위해 그래도 DTO를 쓰자.
func NotificationConfigToDto(src *ent.NotificationConfig) *NotificationConfigDto {
	return &NotificationConfigDto{
		Monday:              src.Monday,
		Tuesday:             src.Tuesday,
		Wednesday:           src.Wednesday,
		Thursday:            src.Thursday,
		Friday:              src.Friday,
		Saturday:            src.Saturday,
		Sunday:              src.Sunday,
		PreferredTimeHour:   src.PreferredTimeHour,
		PreferredTimeMinute: src.PreferredTimeMinute,
		IsActivated:         src.IsActivated,
	}
}
