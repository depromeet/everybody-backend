package dto

type SignUpRequest struct {
	Nickname             string `json:"nickname"`
	NotificationInterval int    `json:"notification_interval"`
	DeviceToken          string `json:"device_token"`
	PushToken            string `json:"push_token"`
	DeviceOS             string `json:"device_os"`
}

type UpdateUserRequest struct {
	Nickname string `json:"nickname"`
}
