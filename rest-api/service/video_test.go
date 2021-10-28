package service

import (
	"testing"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func initializeVideoTest(t *testing.T) *videoService {
	initialize(t)

	return NewVideoService(videoRepo).(*videoService)
}

func TestVideoServiceSave(t *testing.T) {
	t.Run("영상 저장 성공", func(t *testing.T) {
		videoSvc := initializeVideoTest(t)
		expectedVideo := &ent.Video{
			Edges: ent.VideoEdges{
				User:  &ent.User{},
				Album: &ent.Album{},
			},
		}

		videoRepo.On("Save", mock.AnythingOfType("*ent.Video")).Return(expectedVideo, nil)
		video, err := videoSvc.SaveVideo(1, &dto.VideoRequest{})
		assert.NoError(t, err)
		assert.Equal(t, dto.VideoToDto(expectedVideo), video)
	})
}

func TestVideoServiceGetAllByUserID(t *testing.T) {
	t.Run("유저의 전체 영상 조회 성공", func(t *testing.T) {
		videoSvc := initializeVideoTest(t)
		var expectedVideos []*ent.Video

		videoRepo.On("GetAllByUserID", mock.AnythingOfType("int")).Return(expectedVideos, nil)
		videos, err := videoSvc.GetAllVideos(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.VideosToDto(expectedVideos), videos)
	})

	// TODO: error test
}

func TestVideoServiceGet(t *testing.T) {
	t.Run("영상 조회 성공", func(t *testing.T) {
		videoSvc := initializeVideoTest(t)
		expectedVideo := &ent.Video{
			Key: "sample.mp4",
			// dto를 제공하기 위한 URL 맵핑을 하려면 어떤 유저인지를 알아야함.
			Edges: ent.VideoEdges{
				User:  &ent.User{ID: 0},
				Album: &ent.Album{ID: 0},
			},
		}

		videoRepo.On("Get", mock.AnythingOfType("int")).Return(expectedVideo, nil)
		video, err := videoSvc.GetVideo(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.VideoToDto(expectedVideo), video)
	})

	// TODO: error test
}

func TestVideoServiceGets(t *testing.T) {
	t.Run("특정 앨범 사진들 조회 성공", func(t *testing.T) {
		videoSvc := initializeVideoTest(t)
		var expectedVideos []*ent.Video
		videoRepo.On("GetAllByAlbumID", mock.AnythingOfType("int")).Return(expectedVideos, nil)
		videos, err := videoSvc.GetVideos(1)
		assert.NoError(t, err)
		assert.Equal(t, dto.VideosToDto(expectedVideos), videos)
	})

	// TODO: error test
}
