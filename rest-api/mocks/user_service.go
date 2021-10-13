// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	dto "github.com/depromeet/everybody-backend/rest-api/dto"
	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: id
func (_m *UserService) GetUser(id int) (*dto.UserDto, error) {
	ret := _m.Called(id)

	var r0 *dto.UserDto
	if rf, ok := ret.Get(0).(func(int) *dto.UserDto); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignUp provides a mock function with given fields: body
func (_m *UserService) SignUp(body *dto.SignUpRequest) (*dto.UserDto, error) {
	ret := _m.Called(body)

	var r0 *dto.UserDto
	if rf, ok := ret.Get(0).(func(*dto.SignUpRequest) *dto.UserDto); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.SignUpRequest) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}