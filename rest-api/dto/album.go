package dto

import "time"

type AlbumRequest struct {
	// header로 받아오는 걸로?
	// UserID     string `json:"user_id"`
	FolderName string `json:"folder_name"`
}

type AlbumResponse struct {
	FolderName string    `json:"folder_name"`
	CreatedAt  time.Time `json:"created_at"`
}
