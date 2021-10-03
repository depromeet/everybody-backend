package dto

type RegisterDeviceRequest struct {
	DeviceToken string
	PushToken   string
	DeviceOS    string
}
