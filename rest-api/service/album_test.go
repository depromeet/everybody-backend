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
		assert.Equal(t, dto.AlbumToDto(expectedAlbum), newAlbum)
	})

	// TODO: error test
}

func TestAlbumServiceGetAllByUserID(t *testing.T) {
	t.Run("전체 앨범 리스트 조회 성공", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		var expectedAlbums []*ent.Album
		expectedAlbums = append(expectedAlbums, &ent.Album{
			ID:   0,
			Name: "0th",
		})

		albumRepo.On("GetAllByUserID", mock.AnythingOfType("int")).Return(expectedAlbums, nil)
		expectedAlbumsWithPictures := make([]*ent.Album, 0)
		for _, album := range expectedAlbums {
			pictures := make([]*ent.Picture, 0)
			for _, picture := range pictures {
				if album.ID == picture.ID {
					pictures = append(pictures, picture)
				}
			}
			album.Edges.Picture = pictures
			expectedAlbumsWithPictures = append(expectedAlbumsWithPictures, album)
		}

		albums, err := albumSvc.GetAllAlbums(0)
		assert.NoError(t, err)
		assert.Equal(t, dto.AlbumsToDto(expectedAlbumsWithPictures), albums)
	})

	// TODO: error test
}

func TestAlbumServiceGet(t *testing.T) {
	t.Run("각 앨범 조회 성공", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		expectedAlbum := &ent.Album{
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 0},
			},
		}
		var expectedPictures []*ent.Picture

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(expectedAlbum, nil)
		pictureRepo.On("GetAllByAlbumID", mock.AnythingOfType("int")).Return(expectedPictures, nil)
		album, err := albumSvc.GetAlbum(0, 1)
		assert.NoError(t, err)
		assert.Equal(t, dto.AlbumToDto(expectedAlbum), album)
	})

	// TODO: error test
}

func TestAlbumService_UpdateAlbum(t *testing.T) {
	t.Run("본인의 앨범 수정", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		album := &ent.Album{
			ID:   1,
			Name: "기존 이름",
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 1},
			},
		}
		body := &dto.UpdateAlbumRequest{
			Name: "바뀐 이름",
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(album, nil)
		albumRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("*ent.Album")).Run(func(args mock.Arguments) {
			toUpdate := args.Get(1).(*ent.Album)
			album.Name = toUpdate.Name
		}).Return(album, nil)

		updated, err := albumSvc.UpdateAlbum(1, 1, body)
		assert.NoError(t, err)
		assert.Equal(t, updated.Name, "바뀐 이름")
	})

	t.Run("에러) 남의의 앨범 수정", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		album := &ent.Album{
			ID:   1,
			Name: "기존 이름",
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 2},
			},
		}
		body := &dto.UpdateAlbumRequest{
			Name: "바뀐 이름",
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(album, nil)

		updated, err := albumSvc.UpdateAlbum(1, 1, body)
		assert.ErrorIs(t, err, ForbiddenError)
		assert.Nil(t, updated)
	})

	t.Run("에러) 존재하지 않는 앨범 수정", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)

		body := &dto.UpdateAlbumRequest{
			Name: "바뀐 이름",
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(nil, &ent.NotFoundError{})

		updated, err := albumSvc.UpdateAlbum(1, 1, body)
		assert.NotNil(t, err)
		assert.True(t, ent.IsNotFound(err))
		assert.Nil(t, updated)
	})
}

func TestAlbumService_DeleteAlbum(t *testing.T) {
	t.Run("본인의 앨범 삭제", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		album := &ent.Album{
			ID: 1,
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 1},
			},
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(album, nil).Once()
		albumRepo.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

		err := albumSvc.DeleteAlbum(1, 1)
		assert.NoError(t, err)
	})

	t.Run("에러) 남의 앨범 삭제", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		album := &ent.Album{
			ID: 1,
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 2},
			},
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(album, nil).Once()

		err := albumSvc.DeleteAlbum(1, 1)
		assert.ErrorIs(t, err, ForbiddenError)
	})

	t.Run("에러) 존재하지 않는 앨범 삭제", func(t *testing.T) {
		albumSvc := initializeAlbumService(t)
		album := &ent.Album{
			ID: 1,
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 1},
			},
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(album, nil).Once()
		albumRepo.On("Delete", mock.AnythingOfType("int")).Return(&ent.NotFoundError{}).Once()

		err := albumSvc.DeleteAlbum(1, 1)
		assert.True(t, ent.IsNotFound(err))
	})
}
