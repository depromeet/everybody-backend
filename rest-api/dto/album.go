package dto

import "time"

type AlbumRequest struct {
	// header로 받아오는 걸로?
	// UserID     string `json:"user_id"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AlubumsResponse []AlbumResponse

type AlbumResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
