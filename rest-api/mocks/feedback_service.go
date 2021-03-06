// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	dto "github.com/depromeet/everybody-backend/rest-api/dto"
	mock "github.com/stretchr/testify/mock"
)

// FeedbackService is an autogenerated mock type for the FeedbackService type
type FeedbackService struct {
	mock.Mock
}

// SendFeedback provides a mock function with given fields: sender, body
func (_m *FeedbackService) SendFeedback(sender int, body *dto.SendFeedbackRequest) error {
	ret := _m.Called(sender, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *dto.SendFeedbackRequest) error); ok {
		r0 = rf(sender, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
