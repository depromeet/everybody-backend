package service

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type pictureService struct {
	pictureRepo repository.PictureRepositoryInterface
}

type PictureServiceInterface interface {
	SavePicture(userID int, pictureReq *dto.CreatePictureRequest) (*dto.PictureDto, error)
	GetPicture(pictureID int) (*dto.PictureDto, error)
	GetAllPictures(userID int, pictureReq *dto.GetPictureRequest) (dto.PicturesDto, error)
}

func NewPictureService(pictureRepo repository.PictureRepositoryInterface) PictureServiceInterface {
	return &pictureService{
		pictureRepo: pictureRepo,
	}
}

// SavePicture는 API Gateway에서 보낸 picture 정보를(key값 포함) 저장하는 역할
func (s *pictureService) SavePicture(userID int, pictureReq *dto.CreatePictureRequest) (*dto.PictureDto, error) {
	picture := &ent.Picture{
		BodyPart: pictureReq.BodyPart,
		Edges: ent.PictureEdges{
			User:  &ent.User{ID: userID},
			Album: &ent.Album{ID: pictureReq.AlbumID},
		},
		Key: pictureReq.Key,
	}

	p, err := s.pictureRepo.Save(picture)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("사진 저장 완료")
	return dto.PictureToDto(p), nil
}

func (s *pictureService) GetPicture(pictureID int) (*dto.PictureDto, error) {
	picture, err := s.pictureRepo.Get(pictureID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("하나의 사진 조회 완료")
	return dto.PictureToDto(picture), nil
}

// GetAllPictures는 query string으로 오는 uploader, albumID, bodyPart에 맞는 모든 사진들을 조회
func (s *pictureService) GetAllPictures(userID int, pictureReq *dto.GetPictureRequest) (dto.PicturesDto, error) {
	// uploader에 해당하는 사진 조회
	if len(pictureReq.Uploader) > 0 {
		uploaderID, err := strconv.Atoi(pictureReq.Uploader)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		pictures, err := s.pictureRepo.GetAllByUserID(uploaderID)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		log.Info("사용자의 모든 사진들을 조회 완료")
		return dto.PicturesToDto(pictures), nil
	}

	// albumID로 사진 조회
	albumID, err := strconv.Atoi(pictureReq.Album)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// albumID와 bodyPart로 사진 조회
	if len(pictureReq.BodyPart) > 0 {
		pictures, err := s.pictureRepo.FindByAlbumIDAndBodyPart(albumID, pictureReq.BodyPart)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		log.Info("특정 앨범과 신체 부위에 맞는 사진들 조회")
		return dto.PicturesToDto(pictures), nil
	}

	pictures, err := s.pictureRepo.GetAllByAlbumID(albumID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("특정 앨범의 사진들 조회")
	return dto.PicturesToDto(pictures), nil
}
