package handler

import (
	"errors"
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
		return err
	}

	err = ctx.BodyParser(&albumReq)
	if err != nil {
		return err
	}

	newAlbum, err := h.albumService.CreateAlbum(userID, &albumReq)
	if err != nil {
		return err
	}

	return ctx.JSON(dto.AlbumResponse{
		ID:        newAlbum.ID,
		Name:      newAlbum.Name,
		CreatedAt: newAlbum.CreatedAt,
	})
}

func (h *AlbumHandler) GetAllAlbums(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return err
	}

	albums, err := h.albumService.GetAllAlbums(userID)
	if err != nil {
		return err
	}

	albumsResponse := make(dto.AlubumsResponse, 0)
	for _, album := range albums {
		var albumResponse dto.AlbumResponse
		albumResponse.ID = album.ID
		albumResponse.Name = album.Name
		albumResponse.CreatedAt = album.CreatedAt

		albumsResponse = append(albumsResponse, albumResponse)
	}

	return ctx.JSON(albumsResponse)
}

func (h *AlbumHandler) GetAlbum(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "album_id")
	if param == "" {
		return errors.New("params should be provided")
	}

	albumID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	// GetAlbum 에서 각 앨범에 해당하는 pictures도 조회해야 함
	album, _, err := h.albumService.GetAlbum(albumID)
	if err != nil {
		return err
	}

	return ctx.JSON(&dto.AlbumResponse{
		Name: album.Name,
		// PictureList
		CreatedAt: album.CreatedAt,
	})
}
