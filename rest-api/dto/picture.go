package dto

import (
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	log "github.com/sirupsen/logrus"
)

type CreatePictureRequest struct {
	ID       int    `json:"id"`
	AlbumID  int    `json:"album_id"`
	BodyPart string `json:"body_part"`
	// Gateway에서 image key 값도 같이 받음
	Key          string `json:"key"`
	TakenAtYear  int    `json:"taken_at_year"`
	TakenAtMonth int    `json:"taken_at_month"`
	TakenAtDay   int    `json:"taken_at_day"`
}

// 사진 조회할 때 query string으로 오는 것 처리
type GetPictureRequest struct {
	Uploader string `query:"uploader"`
	Album    string `query:"album"`
	BodyPart string `query:"body_part"`
}

// type PictureMultiPart struct {
// 	AlbumID  []string
// 	BodyPart []string
// 	File     []*multipart.FileHeader
// }

type PicturesDto []*PictureDto

type PictureDto struct {
	ID           int    `json:"id"`
	AlbumID      int    `json:"album_id"`
	BodyPart     string `json:"body_part"`
	ThumbnailURL string `json:"thumbnail_url"`
	PreviewURL   string `json:"preview_url"`
	ImageURL     string `json:"image_url"`
	// image의 object key
	Key          string    `json:"key"`
	TakenAtYear  int       `json:"taken_at_year"`
	TakenAtMonth int       `json:"taken_at_month"`
	TakenAtDay   int       `json:"taken_at_day"`
	CreatedAt    time.Time `json:"created_at"`
}

func PictureToDto(src *ent.Picture) *PictureDto {
	thumbnailURL, err := createImageURL(fmt.Sprintf("%d/image/%d/%s", src.Edges.User.ID, 48, src.Key))
	// TODO(umi0410)
	// 일단 이 부분까지 하나 하나 에러처리하긴 번거로울 듯?
	if err != nil {
		log.Error(err)
	}

	previewURL, err := createImageURL(fmt.Sprintf("%d/image/%d/%s", src.Edges.User.ID, 192, src.Key))
	if err != nil {
		log.Error(err)
	}

	imageURL, err := createImageURL(fmt.Sprintf("%d/image/%d/%s", src.Edges.User.ID, 768, src.Key))
	if err != nil {
		log.Error(err)
	}
	year, month, day := util.ConvertTimeToStr(src.TakenAt)
	return &PictureDto{
		ID:           src.ID,
		AlbumID:      src.Edges.Album.ID,
		BodyPart:     src.BodyPart,
		ThumbnailURL: thumbnailURL,
		PreviewURL:   previewURL,
		ImageURL:     imageURL,
		Key:          src.Key,
		TakenAtYear:  year,
		TakenAtMonth: month,
		TakenAtDay:   day,
		CreatedAt:    src.CreatedAt,
	}
}

func PicturesToDto(src []*ent.Picture) PicturesDto {
	picturesDto := make(PicturesDto, 0)

	for _, picture := range src {
		pictureDto := PictureToDto(picture)
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
