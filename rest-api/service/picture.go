package service

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

type pictureService struct {
	pictureRepo repository.PictureRepositoryInterface
	sess        *session.Session
}

type PictureServiceInterface interface {
	SavePicture(userID int, pictureReq *dto.PictureMultiPart) (*dto.PictureDto, error)
	GetAllPictures(userID int) (dto.PicturesDto, error)
	GetPicture(pictureID int) (*dto.PictureDto, error)
}

func NewPictureService(pictureRepo repository.PictureRepositoryInterface, sess *session.Session) PictureServiceInterface {
	return &pictureService{
		pictureRepo: pictureRepo,
		sess:        sess,
	}
}

func (s *pictureService) SavePicture(userID int, pictureReq *dto.PictureMultiPart) (*dto.PictureDto, error) {

	svc := s3.New(s.sess, &aws.Config{
		DisableRestProtocolURICleaning: aws.Bool(true),
		// Region: config.Config.AWS.Region,
	})

	for _, fileHeader := range pictureReq.File {
		f, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		fileBytes, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		// uploader := s3manager.NewUploader(sess)
		output, err := svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(config.Config.AWS.Bucket),
			// TODO: key(object 이름) 값을 어떤 식으로 할 지를 정해야 됨...
			// 임시로 key 값을 filename으로 설정
			Key:         aws.String(fileHeader.Filename),
			Body:        bytes.NewReader(fileBytes),
			ContentType: aws.String(http.DetectContentType(fileBytes)),
		})
		if err != nil {
			return nil, err
		}
		// output 값은 ETag(hash 값) 리턴
		log.Println(output)
	}

	// TODO: 다중 form data 올 경우도 handling 해주어야 할 듯?
	// 현재는 하나씩 온다고 가정하고 구현...
	fileName := pictureReq.File[0].Filename
	bodyPart := pictureReq.BodyPart[0]
	albumID, err := strconv.Atoi(pictureReq.AlbumID[0])
	if err != nil {
		return nil, err
	}

	picture := &ent.Picture{
		BodyPart: bodyPart,
		AlbumID:  albumID,
		Edges: ent.PictureEdges{
			User:  &ent.User{ID: userID},
			Album: &ent.Album{ID: albumID},
		},
		Location: fileName,
	}

	p, err := s.pictureRepo.Save(picture)
	if err != nil {
		return nil, err
	}

	return dto.PictureToDto(p), nil
}

// GetAllPictures는 user의 모든 사진들을 조회
func (s *pictureService) GetAllPictures(userID int) (dto.PicturesDto, error) {
	pictures, err := s.pictureRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	return dto.PicturesToDto(pictures), nil
}

func (s *pictureService) GetPicture(pictureID int) (*dto.PictureDto, error) {
	picture, err := s.pictureRepo.Get(pictureID)
	if err != nil {
		return nil, err
	}

	return dto.PictureToDto(picture), nil
}
