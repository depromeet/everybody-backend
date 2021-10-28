package dto

import (
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type AlbumRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PictureDto는 AlbumID 필드도 가지고 있음
// AlbumDto 할 때 picture에도 album_id를 중복으로
// 가지는 것이 불필요해서 AlbumPicture 구조체를 선언하려고
// 생각하다가 우선은 PictureDto로 하는 걸로...
// type AlbumPicture struct {
// 	PictureID int       `json:"picture_id"`
// 	BodyPart  string    `json:"body_part"`
// 	CreatedAt time.Time `json:"created_at"`
// 	Key  string    `json:"location"`
// }

type AlbumsDto []*AlbumDto
type AlbumDto struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	Pictures  PicturesDto `json:"pictures"`

	// Videos    VideosDto   `json:"videos"`
}

func AlbumsToDto(srcAlbums []*ent.Album) AlbumsDto {
	albumsDto := make(AlbumsDto, 0)

	for _, srcAlbum := range srcAlbums {
		albumDto := AlbumToDto(srcAlbum, srcAlbum.Edges.Picture)
		albumsDto = append(albumsDto, albumDto)
	}

	return albumsDto
}

func AlbumToDto(srcAlbum *ent.Album, srcPictures []*ent.Picture) *AlbumDto {
	picturesDto := PicturesToDto(srcPictures)

	return &AlbumDto{
		ID:        srcAlbum.ID,
		Name:      srcAlbum.Name,
		CreatedAt: srcAlbum.CreatedAt,
		Pictures:  picturesDto,
		// Videos:
	}
}
