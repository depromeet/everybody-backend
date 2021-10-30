package service

import (
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
	SavePicture(userID int, pictureReq *dto.PictureRequest) (*dto.PictureDto, error)
	GetPicture(pictureID int) (*dto.PictureDto, error)
	GetAllPictures(userID int) (dto.PicturesDto, error)
	GetPictures(albumID int, bodyPart string) (dto.PicturesDto, error)
}

func NewPictureService(pictureRepo repository.PictureRepositoryInterface) PictureServiceInterface {
	return &pictureService{
		pictureRepo: pictureRepo,
	}
}

// SavePicture는 API Gateway에서 보낸 picture 정보를(key값 포함) 저장하는 역할
func (s *pictureService) SavePicture(userID int, pictureReq *dto.PictureRequest) (*dto.PictureDto, error) {
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

// GetAllPictures는 user의 모든 사진들을 조회
func (s *pictureService) GetAllPictures(userID int) (dto.PicturesDto, error) {
	pictures, err := s.pictureRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("사용자의 모든 사진들을 조회 완료")
	return dto.PicturesToDto(pictures), nil
}

func (s *pictureService) GetPictures(albumID int, bodyPart string) (dto.PicturesDto, error) {
	// bodyPart가 없다는 것은 특정 앨범 내의 모든 사진들을 조회
	if bodyPart == "" {
		pictures, err := s.pictureRepo.GetAllByAlbumID(albumID)
		if err != nil {
			return nil, err
		}

		log.Info("특정 앨범 내의 모든 사진들 조회 완료")
		return dto.PicturesToDto(pictures), nil
	}

	// albumID와 bodyPart에 맞는 사진들을 조회
	pictures, err := s.pictureRepo.FindByAlbumIDAndBodyPart(albumID, bodyPart)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("특정 앨범 및 신체 부위에 따른 사진들 조회 완료")
	return dto.PicturesToDto(pictures), nil
}
