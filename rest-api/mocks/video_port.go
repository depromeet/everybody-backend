// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// VideoPort is an autogenerated mock type for the VideoPort type
type VideoPort struct {
	mock.Mock
}

// DownloadVideo provides a mock function with given fields: user, imageKeys, duration
func (_m *VideoPort) DownloadVideo(user int, imageKeys []string, duration *float64) (io.Reader, error) {
	ret := _m.Called(user, imageKeys, duration)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(int, []string, *float64) io.Reader); ok {
		r0 = rf(user, imageKeys, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, []string, *float64) error); ok {
		r1 = rf(user, imageKeys, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
