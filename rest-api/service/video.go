package service

import (
	"github.com/depromeet/everybody-backend/rest-api/adapter/video"
	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
)

type videoService struct {
	videoRepo   repository.VideoRepositoryInterface
	pictureRepo repository.PictureRepositoryInterface
	videoPort   video.VideoPort
}

type VideoServiceInterface interface {
	DownloadVideo(userID int, body *dto.DownloadVideoRequest) (io.Reader, error)
	SaveVideo(userID int, videoReq *dto.VideoRequest) (*dto.VideoDto, error)
	GetVideo(videoID int) (*dto.VideoDto, error)
	GetAllVideos(userID int) (dto.VideosDto, error)
}

func NewVideoService(
	videoRepo repository.VideoRepositoryInterface,
	pictureRepo repository.PictureRepositoryInterface,
	videoPort video.VideoPort,
) VideoServiceInterface {
	return &videoService{
		videoRepo:   videoRepo,
		pictureRepo: pictureRepo,
		videoPort:   videoPort,
	}
}

func (s *videoService) DownloadVideo(userID int, body *dto.DownloadVideoRequest) (io.Reader, error) {
	if body.Album == 0 {
		return nil, errors.WithStack(BadRequestError(errors.New("앨범을 설정해주세요.")))
	}
	pictures, err := s.pictureRepo.GetAllByAlbumID(body.Album)
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, picture := range pictures {
		keys = append(keys, picture.Key)
	}

	respBody, err := s.videoPort.DownloadVideo(userID, keys, body.Duration)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (s *videoService) SaveVideo(userID int, videoReq *dto.VideoRequest) (*dto.VideoDto, error) {
	video := &ent.Video{
		Key: videoReq.Key,
		Edges: ent.VideoEdges{
			User: &ent.User{ID: userID},
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
		return nil, errors.WithStack(err)
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
