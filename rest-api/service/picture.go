package service

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type pictureService struct {
	pictureRepo repository.PictureRepositoryInterface
	albumRepo   repository.AlbumRepositoryInterface
}

type PictureServiceInterface interface {
	SavePicture(userID int, pictureReq *dto.CreatePictureRequest) (*dto.PictureDto, error)
	GetPicture(userID, pictureID int) (*dto.PictureDto, error)
	GetAllPictures(userID int, pictureReq *dto.GetPictureRequest) (dto.PicturesDto, error)
	Delete(userID, pictureID int) error
}

func NewPictureService(pictureRepo repository.PictureRepositoryInterface, albumRepo repository.AlbumRepositoryInterface) PictureServiceInterface {
	return &pictureService{
		pictureRepo: pictureRepo,
		albumRepo:   albumRepo,
	}
}

// SavePicture는 API Gateway에서 보낸 picture 정보를(key값 포함) 저장하는 역할
func (s *pictureService) SavePicture(userID int, pictureReq *dto.CreatePictureRequest) (*dto.PictureDto, error) {
	album, err := s.albumRepo.Get(pictureReq.AlbumID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithStack(err)
		}
		// 사진 저장할 앨범 id가 존재하지 않는 등의 에러
		return nil, errors.WithStack(err)
	}

	// 요청한 유저와 앨범의 소유주와 다르다면 그 앨범에 사진이 저장이 되면 안됨
	if userID != album.Edges.User.ID {
		return nil, errors.Wrap(ForbiddenError, "본인의 앨범에만 사진을 업로드할 수 있습니다.")
	}

	takenAt, err := util.ConvertIntToTime(pictureReq.TakenAtYear, pictureReq.TakenAtMonth, pictureReq.TakenAtMonth)
	if err != nil {
		return nil, errors.Wrapf(err, "잘못된 날짜 형식입니다. (%4d, %2d, %2d", pictureReq.TakenAtYear, pictureReq.TakenAtMonth, pictureReq.TakenAtMonth)
	}

	picture := &ent.Picture{
		BodyPart: pictureReq.BodyPart,
		Edges: ent.PictureEdges{
			User:  &ent.User{ID: userID},
			Album: &ent.Album{ID: pictureReq.AlbumID},
		},
		Key:     pictureReq.Key,
		TakenAt: takenAt,
	}

	p, err := s.pictureRepo.Save(picture)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("사진 저장 완료")
	return dto.PictureToDto(p), nil
}

func (s *pictureService) GetPicture(userID, pictureID int) (*dto.PictureDto, error) {
	picture, err := s.pictureRepo.Get(pictureID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if userID != picture.Edges.User.ID {
		return nil, errors.WithStack(errors.New("요청한 유저는 리소스에 접근할 권한이 없습니다."))
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

		// forbiddenerr로 wrapping?
		if userID != uploaderID {
			return nil, errors.Wrap(ForbiddenError, "본인의 사진만을 조회할 수 있습니다.")
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

		if len(pictures) > 0 {
			// forbiddenerr로 wrapping?
			if userID != pictures[0].Edges.User.ID {
				return nil, errors.WithStack(errors.New("요청한 유저는 리소스에 접근할 권한이 없습니다."))
			}
		}

		log.Info("특정 앨범과 신체 부위에 맞는 사진들 조회")
		return dto.PicturesToDto(pictures), nil
	}

	pictures, err := s.pictureRepo.GetAllByAlbumID(albumID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(pictures) > 0 {
		// forbiddenerr로 wrapping?
		if userID != pictures[0].Edges.User.ID {
			return nil, errors.WithStack(errors.New("요청한 유저는 리소스에 접근할 권한이 없습니다."))
		}
	}

	log.Info("특정 앨범의 사진들 조회")
	return dto.PicturesToDto(pictures), nil
}

func (s *pictureService) Delete(userID, pictureID int) error {
	picture, err := s.pictureRepo.Get(pictureID)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.WithMessage(err, "존재하지 않는 사진입니다.")
		}
		return errors.WithMessage(err, "")
	}

	if picture.Edges.User == nil || picture.Edges.User.ID != userID {
		return errors.Wrap(ForbiddenError, "본인의 사진만 삭제할 수 있습니다.")
	}

	err = s.pictureRepo.Delete(pictureID)
	if err != nil {
		return errors.WithMessage(err, "")
	}

	return nil
}
