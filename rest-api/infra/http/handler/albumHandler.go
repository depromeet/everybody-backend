package handler

import (
	"github.com/pkg/errors"
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
)

type AlbumHandler struct {
	albumService service.AlbumServiceInterface
}

func NewAlbumHandler(albumService service.AlbumServiceInterface) *AlbumHandler {
	return &AlbumHandler{
		albumService: albumService,
	}
}

func (h *AlbumHandler) CreateAlbum(ctx *fiber.Ctx) error {
	var albumReq dto.AlbumRequest
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	err = ctx.BodyParser(&albumReq)
	if err != nil {
		return errors.WithStack(err)
	}

	newAlbum, err := h.albumService.CreateAlbum(userID, &albumReq)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(newAlbum)
}

// GetAllAlbums는 album 전체 리스트 조회
func (h *AlbumHandler) GetAllAlbums(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	albums, err := h.albumService.GetAllAlbums(userID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(albums)
}

// GetAlbum은 각 album의 정보와 사진들을 조회
func (h *AlbumHandler) GetAlbum(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "album_id")
	if param == "" {
		return errors.New("params should be provided")
	}

	albumID, err := strconv.Atoi(param)
	if err != nil {
		return errors.WithStack(err)
	}

	// GetAlbum 에서 각 앨범에 해당하는 pictures도 조회해야 함
	albumData, err := h.albumService.GetAlbum(albumID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(albumData)
}
