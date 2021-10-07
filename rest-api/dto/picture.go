package dto

import "time"

type PictureRequest struct {
	ID       int    `json:"id"`
	AlbumID  int    `json:"album_id"`
	BodyPart string `json:"body_part"`
}

type PicturesResponse []PictureResponse

type PictureResponse struct {
	ID        int       `json:"id"`
	AlbumID   int       `json:"album_id"`
	CreatedAt time.Time `json:"created_at"`
	// client한테 어떤 형태로 사진 정보를 줄 지 결정해야함(url, hashed file name 같은...)
}
