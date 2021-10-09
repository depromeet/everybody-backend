package dto

import "github.com/depromeet/everybody-backend/rest-api/ent"

type ConfigureNotificationRequest struct {
	Interval    int  `json:"interval"`
	IsActivated bool `json:"is_activated"`
}

type NotificationConfigDto struct {
	Interval    int  `json:"interval"`
	IsActivated bool `json:"is_activated"`
}

func NotificationConfigToDto(src *ent.NotificationConfig) *NotificationConfigDto {
	return &NotificationConfigDto{
		Interval:    src.Interval,
		IsActivated: src.IsActivated,
	}
}
