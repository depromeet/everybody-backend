package service

import (
	"github.com/pkg/errors"
	"testing"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializePictureTest(t *testing.T) *pictureService {
	initialize(t)

	return NewPictureService(pictureRepo, albumRepo).(*pictureService)
}

func TestPictureServiceSave(t *testing.T) {
	t.Run("사진 업로드, 저장 성공", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		expectedAlbum := &ent.Album{
			Edges: ent.AlbumEdges{
				User: &ent.User{ID: 0},
			},
		}
		expectedPicture := &ent.Picture{
			Edges: ent.PictureEdges{
				User:  &ent.User{ID: 0},
				Album: &ent.Album{ID: 0},
			},
		}

		albumRepo.On("Get", mock.AnythingOfType("int")).Return(expectedAlbum, nil)
		pictureRepo.On("Save", mock.AnythingOfType("*ent.Picture")).Return(expectedPicture, nil)
		picture, err := pictureSvc.SavePicture(0, &dto.CreatePictureRequest{
			ID:       0,
			AlbumID:  0,
			BodyPart: "",
			Key:      "",
		})
		assert.NoError(t, err)
		assert.Equal(t, dto.PictureToDto(expectedPicture), picture)
	})
}

// 사진 전체 조회 테스트 코드
func TestPictureServiceGetAll(t *testing.T) {
	pictureSvc := initializePictureTest(t)
	var expectedPictures []*ent.Picture
	expectedPictures = append(expectedPictures, &ent.Picture{
		BodyPart: "upper",
		TakenAt:  time.Now(),

		Edges: ent.PictureEdges{
			User:  &ent.User{ID: 0},
			Album: &ent.Album{ID: 0},
		}})

	t.Run("유저 전체 사진 조회 성공", func(t *testing.T) {
		pictureRepo.On("GetAllByUserID", mock.AnythingOfType("int")).Return(expectedPictures, nil)
		pictureReq := new(dto.GetPictureRequest)
		pictureReq.Uploader = "0"
		pictures, err := pictureSvc.GetAllPictures(0, pictureReq)
		assert.NoError(t, err)
		assert.Equal(t, dto.PicturesToDto(expectedPictures), pictures)
	})

	t.Run("특정 앨범 내의 사진들 조회 성공", func(t *testing.T) {
		pictureRepo.On("GetAllByAlbumID", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(expectedPictures, nil)
		pictureReq := new(dto.GetPictureRequest)
		pictureReq.Album = "0"
		pictures, err := pictureSvc.GetAllPictures(0, pictureReq)
		assert.NoError(t, err)
		assert.Equal(t, dto.PicturesToDto(expectedPictures), pictures)
	})

	t.Run("특정 앨범 및 신체 부위의 사진들 조회 성공", func(t *testing.T) {
		pictureRepo.On("FindByAlbumIDAndBodyPart", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(expectedPictures, nil)
		pictureReq := new(dto.GetPictureRequest)
		pictureReq.Album = "0"
		pictureReq.BodyPart = "upper"
		pictures, err := pictureSvc.GetAllPictures(0, pictureReq)
		assert.NoError(t, err)
		assert.Equal(t, dto.PicturesToDto(expectedPictures), pictures)
	})

}

func TestPictureServiceGet(t *testing.T) {
	t.Run("사진 조회 성공", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		expectedPicture := &ent.Picture{
			Key: "sample.png",
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.PictureEdges{
				User:  &ent.User{ID: 0},
				Album: &ent.Album{ID: 0},
			},
		}

		pictureRepo.On("Get", mock.AnythingOfType("int")).Return(expectedPicture, nil)
		picture, err := pictureSvc.GetPicture(0, 0)
		assert.NoError(t, err)
		assert.Equal(t, dto.PictureToDto(expectedPicture), picture)
	})

	// TODO: error test
}

func TestPictureService_Delete(t *testing.T) {
	t.Run("성공) 본인 사진 삭제 ", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		picture := &ent.Picture{
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.PictureEdges{
				User:  &ent.User{ID: 1},
				Album: &ent.Album{ID: 1},
			},
		}

		pictureRepo.On("Get", mock.AnythingOfType("int")).Return(picture, nil)
		pictureRepo.On("Delete", mock.AnythingOfType("int")).Return(nil)
		err := pictureSvc.Delete(1, 0)
		assert.NoError(t, err)
	})

	t.Run("에러) 남의 사진 삭제 ", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		picture := &ent.Picture{
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.PictureEdges{
				User:  &ent.User{ID: 2},
				Album: &ent.Album{ID: 1},
			},
		}

		pictureRepo.On("Get", mock.AnythingOfType("int")).Return(picture, nil)
		err := pictureSvc.Delete(1, 0)
		assert.ErrorIs(t, err, ForbiddenError)
	})

	t.Run("에러) 존재하지 않는 사진 삭제 ", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		picture := &ent.Picture{
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.PictureEdges{
				User:  &ent.User{ID: 1},
				Album: &ent.Album{ID: 1},
			},
		}

		pictureRepo.On("Get", mock.AnythingOfType("int")).Return(picture, errors.WithStack(&ent.NotFoundError{}))
		err := pictureSvc.Delete(1, 0)
		assert.True(t, ent.IsNotFound(err))
	})
}
