package dto

type RegisterDeviceRequest struct {
	DeviceToken string `json:"device_token"`
	PushToken   string `json:"push_token"`
	DeviceOS    string `json:"device_os"`
}
