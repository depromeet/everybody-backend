package dto

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type VideoRequest struct {
	Key string `json:"key"`
}

type DownloadVideoRequest struct {
	Album int `json:"album"`
	// TODO: 추후에는 key 자체를 선택해서 한 앨범에서도 따로 따로 다운로드 되면 좋을 듯
	// ImageKeys []string
	Duration *float64 `json:"duration"`
}

type VideoDto struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	VideoURL  string    `json:"video_url"`
}

type VideosDto []*VideoDto

func VideoToDto(srcVideo *ent.Video) *VideoDto {
	// TODO: video-presigned-url 생성 필요
	return &VideoDto{
		ID:        srcVideo.ID,
		Key:       srcVideo.Key,
		CreatedAt: srcVideo.CreatedAt,
	}
}

func VideosToDto(srcVideos []*ent.Video) VideosDto {
	videosDto := make(VideosDto, 0)

	for _, video := range srcVideos {
		videoDto := VideoToDto(video)
		videosDto = append(videosDto, videoDto)
	}

	return videosDto
}
