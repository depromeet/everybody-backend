package dto

import (
	"fmt"
	"time"

	"github.com/depromeet/everybody-backend/rest-api/ent"
)

type AlbumRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AlbumsDto []*AlbumDto
type AlbumDto struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	ThumbnailURL *string   `json:"thumbnail_url"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	// 클라이언트측에서는 부위별로 분류된 사진 리스트가 필요함.
	// TODO: 특정 부위의 사진이 한 장도 없을 때 클라이언트 측에서 맵을 사용하면서 undefined 나 no key(?) 같은 에러를 겪지 않게 하려면
	// 서버측에서 부위에 대한 Enum을 정의해서 각 부위별로 사진이 하나도 없는 부위는 빈 리스트를 전달해줘야할 것 같아요.
	Pictures map[string]PicturesDto `json:"pictures"`
	// Videos    VideosDto   `json:"videos"`
}

var bodyPartKey = []string{"whole", "upper", "lower"}

func AlbumsToDto(srcAlbums []*ent.Album) AlbumsDto {
	albumsDto := make(AlbumsDto, 0)

	for _, srcAlbum := range srcAlbums {
		albumDto := AlbumToDto(srcAlbum)
		albumsDto = append(albumsDto, albumDto)
	}

	return albumsDto
}

func AlbumToDto(srcAlbum *ent.Album) *AlbumDto {
	picturesDto := PicturesToDto(srcAlbum.Edges.Picture)
	picturesMap := make(map[string]PicturesDto)

	// 각 신체부위 key에 대한 초기화
	for _, bodyPart := range bodyPartKey {
		picturesMap[bodyPart] = make(PicturesDto, 0)
	}

	for _, pictureDto := range picturesDto {
		picturesMap[pictureDto.BodyPart] = append(picturesMap[pictureDto.BodyPart], pictureDto)
	}

	// 각 앨범의 대표 썸네일(가장 최신 사진으로)
	var thumbnail *string
	if len(picturesDto) > 0 {
		thumbnail = &picturesDto[len(picturesDto)-1].ThumbnailURL
	}

	duration := time.Now().Sub(srcAlbum.CreatedAt)
	description := fmt.Sprintf("%d일 간의 기록", int(duration.Hours())/24+1)

	return &AlbumDto{
		ID:           srcAlbum.ID,
		Name:         srcAlbum.Name,
		ThumbnailURL: thumbnail,
		Description:  description,
		CreatedAt:    srcAlbum.CreatedAt,
		Pictures:     picturesMap,
	}
}

// 현재 기획상으로는 앨범의 이름만 변경할 수 있음.
type UpdateAlbumRequest struct {
	Name string
}
