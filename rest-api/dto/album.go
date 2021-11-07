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
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	// 클라이언트측에서는 부위별로 분류된 사진 리스트가 필요함.
	// TODO: 특정 부위의 사진이 한 장도 없을 때 클라이언트 측에서 맵을 사용하면서 undefined 나 no key(?) 같은 에러를 겪지 않게 하려면
	// 서버측에서 부위에 대한 Enum을 정의해서 각 부위별로 사진이 하나도 없는 부위는 빈 리스트를 전달해줘야할 것 같아요.
	Pictures map[string]PicturesDto `json:"pictures"`

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
	picturesMap := make(map[string]PicturesDto)
	for _, pictureDto := range picturesDto {
		picturesMap[pictureDto.BodyPart] = append(picturesMap[pictureDto.BodyPart], pictureDto)
	}
	return &AlbumDto{
		ID:        srcAlbum.ID,
		Name:      srcAlbum.Name,
		CreatedAt: srcAlbum.CreatedAt,
		Pictures:  picturesMap,
		// Videos:
	}
}
