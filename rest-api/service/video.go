package service

import (
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type videoService struct {
	videoRepo repository.VideoRepositoryInterface
}

type VideoServiceInterface interface {
	SaveVideo(userID int, videoReq *dto.VideoRequest) (*dto.VideoDto, error)
	GetVideo(videoID int) (*dto.VideoDto, error)
	GetAllVideos(userID int) (dto.VideosDto, error)
	GetVideos(albumID int) (dto.VideosDto, error)
}

func NewVideoService(videoRepo repository.VideoRepositoryInterface) VideoServiceInterface {
	return &videoService{
		videoRepo: videoRepo,
	}
}

func (s *videoService) SaveVideo(userID int, videoReq *dto.VideoRequest) (*dto.VideoDto, error) {
	video := &ent.Video{
		Key: videoReq.Key,
		Edges: ent.VideoEdges{
			User:  &ent.User{ID: userID},
			Album: &ent.Album{ID: videoReq.AlbumID},
		},
	}

	v, err := s.videoRepo.Save(video)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.VideoToDto(v), nil
}

// GetVideo는 영상 하나 조회
func (s *videoService) GetVideo(videoID int) (*dto.VideoDto, error) {
	v, err := s.videoRepo.Get(videoID)
	if err != nil {
		return nil, err
	}

	log.Info("영상 조회 완료")
	return dto.VideoToDto(v), nil
}

// GetAllVideos는 유저의 전체 영상 조회
func (s *videoService) GetAllVideos(userID int) (dto.VideosDto, error) {
	videos, err := s.videoRepo.GetAllByUserID(userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("유저의 전체 영상 조회 완료")
	return dto.VideosToDto(videos), nil
}

// GetVideos는 앨범의 영상들을 조회
func (s *videoService) GetVideos(albumID int) (dto.VideosDto, error) {
	videos, err := s.videoRepo.GetAllByAlbumID(albumID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("앨범의 영상 조회 완료")
	return dto.VideosToDto(videos), nil
}
