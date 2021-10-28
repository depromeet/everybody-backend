package dto

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type VideoRequest struct {
	AlbumID int    `json:"album_id"`
	Key     string `json:"key"`
}

type VideoDto struct {
	ID        int       `json:"id"`
	AlbumID   int       `json:"album_id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
	VideoURL  string    `json:"video_url"`
}

type VideosDto []*VideoDto

func VideoToDto(srcVideo *ent.Video) *VideoDto {
	// TODO: video-presigned-url 생성 필요
	return &VideoDto{
		ID:        srcVideo.ID,
		AlbumID:   srcVideo.Edges.Album.ID,
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
