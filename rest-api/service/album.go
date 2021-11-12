package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type albumService struct {
	albumRepo   repository.AlbumRepositoryInterface
	pictureRepo repository.PictureRepositoryInterface
}

type AlbumServiceInterface interface {
	CreateAlbum(userID int, albumReq *dto.AlbumRequest) (*dto.AlbumDto, error)
	GetAllAlbums(userID int) (dto.AlbumsDto, error)
	GetAlbum(userID, albumID int) (*dto.AlbumDto, error)
	UpdateAlbum(userID, albumID int, body *dto.UpdateAlbumRequest) (*dto.AlbumDto, error)
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
		return nil, errors.WithStack(err)
	}

	log.Info("앨범 생성 완료")
	return dto.AlbumToDto(newAlbum), nil
}

// GetAllAlbums는 album의 전체 리스트를 조회
func (s *albumService) GetAllAlbums(userID int) (dto.AlbumsDto, error) {
	albums, err := s.albumRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(albums) == 0 {
		return nil, errors.WithStack(errors.New("해당하는 리소스를 찾지 못했습니다."))
	}

	log.Info("전체 앨범 조회 완료 (그 속 사진들은 Eager loading으로 조회)")
	return dto.AlbumsToDto(albums), nil
}

// GetAlbum은 alubm 정보와 album의 사진 정보들도 조회
func (s *albumService) GetAlbum(userID, albumID int) (*dto.AlbumDto, error) {
	album, err := s.albumRepo.Get(albumID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if userID != album.Edges.User.ID {
		return nil, errors.WithStack(errors.New("요청한 유저는 리소스에 접근할 권한이 없습니다."))
	}

	log.Info("앨범과 그 앨범의 사진들을 조회 완료")
	return dto.AlbumToDto(album), err
}

func (s *albumService) UpdateAlbum(userID, albumID int, body *dto.UpdateAlbumRequest) (*dto.AlbumDto, error) {
	album, err := s.albumRepo.Get(albumID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.Wrap(err, "존재하지 않는 앨범입니다.")
		}

		return nil, errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}
	if album.Edges.User == nil || album.Edges.User.ID != userID {
		return nil, errors.Wrap(ForbiddenError, "본인의 앨범만 수정할 수 있습니다.")
	}

	updated, err := s.albumRepo.Update(albumID, &ent.Album{
		Name: body.Name,
	})
	if err != nil {
		return nil, errors.WithMessage(err, err.Error())
	}
	log.Infof("앨범 수정 완료: %v", updated)

	return dto.AlbumToDto(updated), nil
}
