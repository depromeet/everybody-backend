// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ent "github.com/depromeet/everybody-backend/rest-api/ent"
	mock "github.com/stretchr/testify/mock"
)

// PictureRepositoryInterface is an autogenerated mock type for the PictureRepositoryInterface type
type PictureRepositoryInterface struct {
	mock.Mock
}

// FindByAlbumIDAndBodyPart provides a mock function with given fields: albumID, bodyPart
func (_m *PictureRepositoryInterface) FindByAlbumIDAndBodyPart(albumID int, bodyPart string) ([]*ent.Picture, error) {
	ret := _m.Called(albumID, bodyPart)

	var r0 []*ent.Picture
	if rf, ok := ret.Get(0).(func(int, string) []*ent.Picture); ok {
		r0 = rf(albumID, bodyPart)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Picture)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(albumID, bodyPart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: pictureID
func (_m *PictureRepositoryInterface) Get(pictureID int) (*ent.Picture, error) {
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

// GetAllByAlbumID provides a mock function with given fields: albumID
func (_m *PictureRepositoryInterface) GetAllByAlbumID(albumID int) ([]*ent.Picture, error) {
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

// GetAllByUserID provides a mock function with given fields: userID
func (_m *PictureRepositoryInterface) GetAllByUserID(userID int) ([]*ent.Picture, error) {
	ret := _m.Called(userID)

	var r0 []*ent.Picture
	if rf, ok := ret.Get(0).(func(int) []*ent.Picture); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.Picture)
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

// Save provides a mock function with given fields: picture
func (_m *PictureRepositoryInterface) Save(picture *ent.Picture) (*ent.Picture, error) {
	ret := _m.Called(picture)

	var r0 *ent.Picture
	if rf, ok := ret.Get(0).(func(*ent.Picture) *ent.Picture); ok {
		r0 = rf(picture)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.Picture)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ent.Picture) error); ok {
		r1 = rf(picture)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
