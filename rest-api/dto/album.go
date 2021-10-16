package dto

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type AlbumRequest struct {
	// header로 받아오는 걸로?
	// UserID     string `json:"user_id"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AlbumPicture struct {
	PictureID int       `json:"picture_id"`
	BodyPart  string    `json:"body_part"`
	CreatedAt time.Time `json:"created_at"`
	Location  string    `json:"location"`
}

type AlbumsDto []AlbumDto
type AlbumDto struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	Pictures  []AlbumPicture `json:"pictures"`
}

func AlbumsToDto(src []*ent.Album) AlbumsDto {
	albumsDto := make(AlbumsDto, 0)
	for _, srcAlbum := range src {
		albumDto := AlbumDto{
			ID:        srcAlbum.ID,
			Name:      srcAlbum.Name,
			CreatedAt: srcAlbum.CreatedAt,
		}

		albumsDto = append(albumsDto, albumDto)
	}

	return albumsDto
}

func AlbumToDto(src *ent.Album, srcPictures []*ent.Picture) *AlbumDto {
	picturesDto := make([]AlbumPicture, 0)
	if srcPictures != nil {
		for _, srcPicture := range srcPictures {
			pictureDto := AlbumPicture{
				PictureID: srcPicture.ID,
				BodyPart:  srcPicture.BodyPart,
				CreatedAt: srcPicture.CreatedAt,
				Location:  srcPicture.Location,
			}

			picturesDto = append(picturesDto, pictureDto)
		}
	}

	return &AlbumDto{
		ID:        src.ID,
		Name:      src.Name,
		CreatedAt: src.CreatedAt,
		Pictures:  picturesDto,
	}
}
