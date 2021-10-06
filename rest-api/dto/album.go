package dto

import "time"

type AlbumRequest struct {
	// header로 받아오는 걸로?
	// UserID     string `json:"user_id"`
	ID         int    `json:"id"`
	FolderName string `json:"folder_name"`
}

type AlbumResponse struct {
	ID         int       `json:"id,omitempty"`
	FolderName string    `json:"folder_name,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
