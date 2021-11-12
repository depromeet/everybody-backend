package handler

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	param := util.GetParams(ctx, "album_id")
	if param == "" {
		return errors.New("params should be provided")
	}

	albumID, err := strconv.Atoi(param)
	if err != nil {
		return errors.WithStack(err)
	}

	// GetAlbum 에서 각 앨범에 해당하는 pictures도 조회해야 함
	albumData, err := h.albumService.GetAlbum(userID, albumID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(albumData)
}

func (h *AlbumHandler) UpdateAlbum(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	albumID, err := strconv.Atoi(ctx.Params("album_id"))
	if err != nil {
		return errors.WithStack(err)
	}

	body := new(dto.UpdateAlbumRequest)
	err = ctx.BodyParser(body)
	if err != nil {
		return errors.WithStack(err)
	}

	album, err := h.albumService.UpdateAlbum(userID, albumID, body)
	if err != nil {
		return errors.WithMessage(err, "")
	}

	return ctx.JSON(album)
}

func (h *AlbumHandler) DeleteAlbum(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	albumID, err := strconv.Atoi(ctx.Params("album_id"))
	if err != nil {
		return errors.WithStack(err)
	}

	err = h.albumService.DeleteAlbum(userID, albumID)
	if err != nil {
		return errors.WithMessage(err, "")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
