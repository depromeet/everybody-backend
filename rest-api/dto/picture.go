package dto

import (
	"mime/multipart"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type PictureRequest struct {
	ID       int    `json:"id"`
	AlbumID  int    `json:"album_id"`
	BodyPart string `json:"body_part"`
}

type PictureMultiPart struct {
	AlbumID  []string
	BodyPart []string
	File     []*multipart.FileHeader
}

type PicturesDto []PictureDto

type PictureDto struct {
	ID        int       `json:"id"`
	AlbumID   int       `json:"album_id"`
	BodyPart  string    `json:"body_part"`
	CreatedAt time.Time `json:"created_at"`
	// client한테 어떤 형태로 사진 정보를 줄 지 결정해야함(url, hashed file name 같은...)
	Location string `json:"location"`
}

func PictureToDto(src *ent.Picture) *PictureDto {
	return &PictureDto{
		ID:        src.ID,
		AlbumID:   src.AlbumID,
		BodyPart:  src.BodyPart,
		CreatedAt: src.CreatedAt,
		Location:  src.Location,
	}
}

func PicturesToDto(src []*ent.Picture) PicturesDto {
	picturesDto := make(PicturesDto, 0)

	for _, picture := range src {
		pictureDto := PictureDto{
			ID:        picture.ID,
			AlbumID:   picture.AlbumID,
			BodyPart:  picture.BodyPart,
			CreatedAt: picture.CreatedAt,
			Location:  picture.Location,
		}

		picturesDto = append(picturesDto, pictureDto)
	}

	return picturesDto
}
