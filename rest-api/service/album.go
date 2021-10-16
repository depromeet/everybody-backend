package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
)

type albumService struct {
	albumRepo   repository.AlbumRepositoryInterface
	pictureRepo repository.PictureRepositoryInterface
}

type AlbumServiceInterface interface {
	CreateAlbum(userID int, albumReq *dto.AlbumRequest) (*dto.AlbumDto, error)
	GetAllAlbums(userID int) (dto.AlbumsDto, error)
	GetAlbum(albumID int) (*dto.AlbumDto, error)
}

func NewAlbumService(albumRepo repository.AlbumRepositoryInterface, pictureRepo repository.PictureRepositoryInterface) AlbumServiceInterface {
	return &albumService{
		albumRepo:   albumRepo,
		pictureRepo: pictureRepo,
	}
}

func (s *albumService) CreateAlbum(userID int, albumReq *dto.AlbumRequest) (*dto.AlbumDto, error) {
	album := &ent.Album{
		Name: albumReq.Name,
		Edges: ent.AlbumEdges{
			User: &ent.User{ID: userID},
		},
	}

	newAlbum, err := s.albumRepo.Create(album)
	if err != nil {
		return nil, err
	}

	return dto.AlbumToDto(newAlbum, nil), nil
}

// GetAllAlbums는 album의 전체 리스트를 조회(사진 정보는 조회X)
func (s *albumService) GetAllAlbums(userID int) (dto.AlbumsDto, error) {
	albums, err := s.albumRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, err
	}

	// album 각각의 사진들도 조회?

	return dto.AlbumsToDto(albums), nil
}

// GetAlbum은 alubm 정보와 album의 사진 정보들도 조회
func (s *albumService) GetAlbum(albumID int) (*dto.AlbumDto, error) {
	album, err := s.albumRepo.Get(albumID)
	if err != nil {
		return nil, err
	}

	// albumID에 해당하는 pictrues 목록 조회 기능 필요
	pictures, err := s.pictureRepo.GetAllByAlbumID(albumID)
	if err != nil {
		return nil, err
	}

	return dto.AlbumToDto(album, pictures), err
}
