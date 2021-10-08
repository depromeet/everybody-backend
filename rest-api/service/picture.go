package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
)

type pictureService struct {
	pictureRepo repository.PictureRepositoryInterface
}

type PictureServiceInterface interface {
	SavePicture(pictureReq *dto.PictureRequest) (bool, error)
	GetAllPictures(albumID int) ([]*ent.Picture, error)
	GetPicture(pictureID int) (*ent.Picture, error)
}

func NewPictureService(pictureRepo repository.PictureRepositoryInterface) PictureServiceInterface {
	return &pictureService{
		pictureRepo: pictureRepo,
	}
}

func (s *pictureService) SavePicture(pictureReq *dto.PictureRequest) (bool, error) {
	picture := &ent.Picture{
		BodyPart: pictureReq.BodyPart,
		Edges: ent.PictureEdges{
			Album: &ent.Album{ID: pictureReq.AlbumID},
		},
	}

	err := s.pictureRepo.Save(picture)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *pictureService) GetAllPictures(albumID int) ([]*ent.Picture, error) {
	pictures, err := s.pictureRepo.GetAllByAlbumID(albumID)
	if err != nil {
		return nil, err
	}

	return pictures, nil
}

func (s *pictureService) GetPicture(pictureID int) (*ent.Picture, error) {
	picture, err := s.pictureRepo.Get(pictureID)
	if err != nil {
		return nil, err
	}

	return picture, nil
}
