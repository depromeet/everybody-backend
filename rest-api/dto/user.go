package dto

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
