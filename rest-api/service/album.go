package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
)

type albumService struct {
	albumRepo repository.AlbumRepositoryInterface
}

type AlbumServiceInterface interface {
	CreateAlbum(userID string, albumReq *dto.AlbumRequest) (*ent.Album, error)
	GetAlbum(albumID int) (*ent.Album, error)
}

func NewAlbumService(albumRepo repository.AlbumRepositoryInterface) AlbumServiceInterface {
	return &albumService{
		albumRepo: albumRepo,
	}
}

func (s *albumService) CreateAlbum(userID string, albumReq *dto.AlbumRequest) (*ent.Album, error) {
	album := &ent.Album{
		FolderName: albumReq.FolderName,
		Edges: ent.AlbumEdges{
			User: &ent.User{ID: userID},
		},
	}

	newAlbum, err := s.albumRepo.Create(album)
	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func (s *albumService) GetAlbum(albumID int) (*ent.Album, error) {
	albumData, err := s.albumRepo.Get(albumID)
	if err != nil {
		return nil, err
	}

	return albumData, err
}
