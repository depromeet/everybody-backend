package service

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/ent"
	entUser "github.com/depromeet/everybody-backend/rest-api/ent/user"
	"github.com/depromeet/everybody-backend/rest-api/repository"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	ErrDuplicatedUserID          = errors.New("고유한 유저 ID 생성에 실패했습니다.")
	ErrMissingNotificationConfig = errors.New("알림 설정 관련 정보가 없습니다.")
	ErrMissingDevice             = errors.New("디바이스 정보가 없습니다.")
	signUpDefaultNickname        = "끈육몬"
	randomNicknameAdjectives     = []string{
		"배고픈", "귀여운", "열정적인", "빛나는", "튼튼한", "힘찬", "고독한", "따뜻한", "깜찍한", "꾸준한", "단단한",
	}
	randomNicknameNouns = []string{
		"고양이", "멍멍이", "댕댕이", "냥이", "기요밍", "튼튼이", "헬린이", "뽀로로", "뿌앙이",
	}
	// randomProfileImageKey 는 public drive s3에 미리 업로드 해둔 샘플이다.
	randomProfileImageKey = []string{
		"beam-1.png", "beam-2.png", "beam-3.png", "beam-4.png", "beam-5.png", "beam-6.png",
	}
	defaultMotto     = "천천히 그리고 꾸준히!"
	defaultAlbumName = "눈바디"
)

type UserService interface {
	SignUp(body *dto.SignUpRequest) (*dto.UserDto, error)
	GetUser(id int) (*dto.UserDto, error)
	UpdateUser(id int, body *dto.UpdateUserRequest) (*dto.UserDto, error)
	UpdateProfileImage(id int, body *dto.UpdateProfileImageRequest) (*dto.UserDto, error)
}

func NewUserService(
	userRepo repository.UserRepository,
	notificationService NotificationService,
	deviceService DeviceService,
	albumService AlbumServiceInterface) UserService {
	return &userService{
		userRepo:            userRepo,
		notificationService: notificationService,
		deviceService:       deviceService,
		albumService:        albumService,
	}
}

// UserService 의 구현체
type userService struct {
	userRepo            repository.UserRepository
	notificationService NotificationService
	deviceService       DeviceService
	albumService        AlbumServiceInterface
}

// SignUp 는 유저 생성 후 해당 유저의 Device를 등록합니다.
// 유저의 ID는 랜덤하고 고유한 UUID 입니다.
// 닉네임은 정의되지 않은 경우 "끈육몬"이 됨.
// TODO: 트랜잭션 롤백이 안됨. 유저를 만들고 다른 것들을 만들다가 종료되면..?
func (s *userService) SignUp(body *dto.SignUpRequest) (*dto.UserDto, error) {
	user, err := s.userRepo.Create(&ent.User{
		Nickname:     s.generateUniqueNickname(),
		Motto:        defaultMotto,
		Height:       body.Height,
		Weight:       body.Weight,
		ProfileImage: s.getRandomProfileImageKey(),
		Kind:         entUser.Kind(body.Kind),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Infof("유저를 생성했습니다. User(id=%d)", user.ID)

	if body.NotificationConfig == nil {
		log.Info("NotificationConfig가 존재하지 않습니다. Default 값으로 설정합니다.")
		body.NotificationConfig = s.defaultConfigureNotificationRequest()
	}
	_, err = s.notificationService.Configure(user.ID, body.NotificationConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Infof("알림 설정을 생성했습니다. NotificaitonConfig(user.id=%d)", user.ID)

	if body.Device == nil {
		return nil, errors.WithStack(ErrMissingDevice)
	}
	device, err := s.deviceService.Register(user.ID, body.Device)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Infof("디바이스 정보를 생성했습니다. Device(id=%d)", device.ID)

	_, err = s.albumService.CreateAlbum(user.ID, &dto.AlbumRequest{
		Name: defaultAlbumName,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	log.Info("유저의 default 앨범을 생성했습니다.")

	return dto.UserToDto(user), nil
}

// GetUser 는 유저 정보를 조회합니다.
func (s *userService) GetUser(id int) (*dto.UserDto, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.UserToDto(user), nil
}

// GetUser 는 유저가 수정할 수 있는 유저 정보의 필드들을 수정합니다.
func (s *userService) UpdateUser(id int, body *dto.UpdateUserRequest) (*dto.UserDto, error) {
	user, err := s.userRepo.Update(id, &ent.User{
		Nickname: body.Nickname,
		Motto:    body.Motto,
		Height:   body.Height,
		Weight:   body.Weight,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.UserToDto(user), nil
}

func (s *userService) UpdateProfileImage(id int, body *dto.UpdateProfileImageRequest) (*dto.UserDto, error) {
	user, err := s.userRepo.UpdateProfileImage(id, body.ProfileImage)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dto.UserToDto(user), nil
}

func (s *userService) defaultConfigureNotificationRequest() *dto.ConfigureNotificationRequest {
	return &dto.ConfigureNotificationRequest{
		Monday:              true,
		Tuesday:             false,
		Wednesday:           true,
		Thursday:            false,
		Friday:              true,
		Saturday:            false,
		Sunday:              false,
		PreferredTimeHour:   20,
		PreferredTimeMinute: 0,
		IsActivated:         true,
	}
}
func (s *userService) getRandomProfileImageKey() string {
	return randomProfileImageKey[rand.Intn(len(randomProfileImageKey))]
}
func (s *userService) generateUniqueNickname() string {
	suffix := 0
	adj := randomNicknameAdjectives[rand.Intn(len(randomNicknameAdjectives))]
	noun := randomNicknameNouns[rand.Intn(len(randomNicknameNouns))]

	last, err := s.userRepo.FindByNicknameContainingOrderByNicknameDesc(adj + noun)
	if ent.IsNotFound(err) {
		return adj + noun
	}
	if len(adj+noun) != len(last.Nickname) {
		suffix, err = strconv.Atoi(last.Nickname[len(adj+noun):])

		if err != nil {
			log.Errorf("랜덤 닉네임 생성 도중 오류가 발생했습니다. Suffix가 숫자가 아닙니다: %s", last.Nickname)
			panic(errors.Wrap(err, "랜덤 닉네임 생성 도중 오류가 발생했습니다. Suffix가 숫자가 아닙니다."))
		}
	}

	return fmt.Sprintf("%s%s%d", adj, noun, suffix+1)
}
