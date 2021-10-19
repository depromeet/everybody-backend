package service

import (
	"testing"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializeAlbumService(t *testing.T) *albumService {
	initialize(t)

	return NewAlbumService(albumRepo, pictureRepo).(*albumService)
}

func TestAlbumServiceCreate(t *testing.T) {
	t.Run("앨범 생성 성공", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		expectedAlbum := new(ent.Album)

		albumRepo.On("Create", mock.AnythingOfType("*ent.Album")).Return(expectedAlbum, nil)
		newAlbum, err := albumSvc.CreateAlbum(1, &dto.AlbumRequest{})
		assert.NoError(t, err)
		assert.Equal(t, dto.AlbumToDto(expectedAlbum, nil), newAlbum)
	})

	// TODO: error test
}

func TestAlbumServiceGetAllByUserID(t *testing.T) {
	t.Run("전체 앨범 리스트 조회 성공", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		var expectedAlbums []*ent.Album

		albumRepo.On("GetAllByUserID", mock.AnythingOfType("int")).Return(expectedAlbums, nil)
		albums, err := albumSvc.GetAllAlbums(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.AlbumsToDto(expectedAlbums), albums)
	})

	// TODO: error test
}

func TestAlbumServiceGet(t *testing.T) {
	t.Run("각 앨범 조회 성공", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		expectedAlbum := new(ent.Album)
		var expectedPictures []*ent.Picture

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(expectedAlbum, nil)
		pictureRepo.On("GetAllByAlbumID", mock.AnythingOfType("int")).Return(expectedPictures, nil)
		album, err := albumSvc.GetAlbum(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.AlbumToDto(expectedAlbum, expectedPictures), album)
	})

	// TODO: error test
}