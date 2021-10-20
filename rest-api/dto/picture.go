package dto

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"mime/multipart"
	"strings"
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

func createImageURL(imageKey string) (string, error) {
	// TODO(umi0410): Signed URL을 이용한 맵핑 최적화
	// 사실 사진 한 장 한 장의 주소들을 맵핑할 때마다 signed url을 재생성하는 게 조금
	// 비효율적일 것 같긴한데 일단은 간단하게 구현하느라
	// 이렇게 매번 새로운 Signed URL을 생성하도록 해놨음.
	// 원래는 어차피 다 같은 Signature로 제공할 수 있어서
	// 이렇게 매번 생성할 필요는 없음.
	pk, err := sign.LoadPEMPrivKey(strings.NewReader(config.Config.ImagePrivateKey))
	if err != nil {
		return "", err
	}

	signer := sign.NewURLSigner(config.Config.ImagePublicKeyID, pk)
	url, err := signer.Sign(fmt.Sprintf("%s/%s", config.Config.ImageRootUrl, imageKey), time.Now().Add(time.Minute))
	if err != nil {
		return "", err
	}

	return url, nil
}
