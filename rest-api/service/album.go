package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
)

type albumService struct {
	albumRepo repository.AlbumRepositoryInterface
	// pictureRepo repository.PictureRepositoryInterface
}

type AlbumServiceInterface interface {
	CreateAlbum(userID string, albumReq *dto.AlbumRequest) (*ent.Album, error)
	GetAllAlbums(userID string) ([]*ent.Album, error)
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

func (s *albumService) GetAllAlbums(userID string) ([]*ent.Album, error) {
	albums, err := s.albumRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// GetAlbum은 alubm 정보와 album의 사진 정보들도 조회
func (s *albumService) GetAlbum(albumID int) (*ent.Album, error) {
	albumData, err := s.albumRepo.Get(albumID)
	if err != nil {
		return nil, err
	}

	// albumID에 해당하는 pictrues 목록 조회 기능 필요

	return albumData, err
}
