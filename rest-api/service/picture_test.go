package service

import (
	"testing"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializePictureTest(t *testing.T) *pictureService {
	initialize(t)

	return NewPictureService(pictureRepo, nil).(*pictureService)
}

// TODO: AWS와 session 연결이랑, multipart 부분 test 필요
// func TestPictureServiceSave(t *testing.T) {
// 	t.Run("사진 업로드, 저장 성공", func(t *testing.T) {
// 		expected := &ent.Picture{}
// 		pictureSvc := initializePictureTest(t)
// 		mockPicture := new(ent.Picture)

// 		_, _ = pictureSvc.SavePicture(1, &dto.PictureMultiPart{})
// 		pictureRepo.On("Save", mock.AnythingOfType("*ent.Picture")).Return(mockPicture, nil)

// 		assert.Equal(t, expected, mockPicture)
// 	})
// }

func TestPictureServiceGetAllByUserID(t *testing.T) {
	t.Run("유저의 전체 사진들 조회 성공", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		var expectedPictures []*ent.Picture

		pictureRepo.On("GetAllByUserID", mock.AnythingOfType("int")).Return(expectedPictures, nil)
		pictures, err := pictureSvc.GetAllPictures(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.PicturesToDto(expectedPictures), pictures)
	})

	// TODO: error test
}

func TestPictureServiceGet(t *testing.T) {
	t.Run("사진 조회 성공", func(t *testing.T) {
		pictureSvc := initializePictureTest(t)
		expectedPicture := &ent.Picture{
			Location: "sample.png",
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.PictureEdges{
				User: &ent.User{ID: 0},
			},
		}

		pictureRepo.On("Get", mock.AnythingOfType("int")).Return(expectedPicture, nil)
		picture, err := pictureSvc.GetPicture(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.PictureToDto(expectedPicture), picture)
	})

	// TODO: error test
}
