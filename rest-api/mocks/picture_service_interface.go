// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	dto "github.com/depromeet/everybody-backend/rest-api/dto"
	ent "github.com/depromeet/everybody-backend/rest-api/ent"

	mock "github.com/stretchr/testify/mock"
)

// PictureServiceInterface is an autogenerated mock type for the PictureServiceInterface type
type PictureServiceInterface struct {
	mock.Mock
}

// GetAllPictures provides a mock function with given fields: albumID
func (_m *PictureServiceInterface) GetAllPictures(albumID int) ([]*ent.Picture, error) {
	ret := _m.Called(albumID)

	var r0 []*ent.Picture
	if rf, ok := ret.Get(0).(func(int) []*ent.Picture); ok {
		r0 = rf(albumID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Picture)
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

// GetPicture provides a mock function with given fields: pictureID
func (_m *PictureServiceInterface) GetPicture(pictureID int) (*ent.Picture, error) {
	ret := _m.Called(pictureID)

	var r0 *ent.Picture
	if rf, ok := ret.Get(0).(func(int) *ent.Picture); ok {
		r0 = rf(pictureID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Picture)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(pictureID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePicture provides a mock function with given fields: pictureReq
func (_m *PictureServiceInterface) SavePicture(pictureReq *dto.PictureRequest) (bool, error) {
	ret := _m.Called(pictureReq)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*dto.PictureRequest) bool); ok {
		r0 = rf(pictureReq)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*dto.PictureRequest) error); ok {
		r1 = rf(pictureReq)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}