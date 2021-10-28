// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	dto "github.com/depromeet/everybody-backend/rest-api/dto"
	mock "github.com/stretchr/testify/mock"
)

// VideoServiceInterface is an autogenerated mock type for the VideoServiceInterface type
type VideoServiceInterface struct {
	mock.Mock
}

// GetAllVideos provides a mock function with given fields: userID
func (_m *VideoServiceInterface) GetAllVideos(userID int) (dto.VideosDto, error) {
	ret := _m.Called(userID)

	var r0 dto.VideosDto
	if rf, ok := ret.Get(0).(func(int) dto.VideosDto); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dto.VideosDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideo provides a mock function with given fields: videoID
func (_m *VideoServiceInterface) GetVideo(videoID int) (*dto.VideoDto, error) {
	ret := _m.Called(videoID)

	var r0 *dto.VideoDto
	if rf, ok := ret.Get(0).(func(int) *dto.VideoDto); ok {
		r0 = rf(videoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.VideoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(videoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVideos provides a mock function with given fields: albumID
func (_m *VideoServiceInterface) GetVideos(albumID int) (dto.VideosDto, error) {
	ret := _m.Called(albumID)

	var r0 dto.VideosDto
	if rf, ok := ret.Get(0).(func(int) dto.VideosDto); ok {
		r0 = rf(albumID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dto.VideosDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(albumID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveVideo provides a mock function with given fields: userID, videoReq
func (_m *VideoServiceInterface) SaveVideo(userID int, videoReq *dto.VideoRequest) (*dto.VideoDto, error) {
	ret := _m.Called(userID, videoReq)

	var r0 *dto.VideoDto
	if rf, ok := ret.Get(0).(func(int, *dto.VideoRequest) *dto.VideoDto); ok {
		r0 = rf(userID, videoReq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.VideoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *dto.VideoRequest) error); ok {
		r1 = rf(userID, videoReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
