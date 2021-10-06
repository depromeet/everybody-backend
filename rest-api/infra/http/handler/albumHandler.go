package handler

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
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
	userID := util.GetRequestUserID(ctx)
	err := ctx.BodyParser(&albumReq)
	if err != nil {
		log.Printf("invalid json body: %v", err)
		return ctx.JSON(err)
	}

	newAlbum, err := h.albumService.CreateAlbum(userID, &albumReq)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(dto.AlbumResponse{
		ID:         newAlbum.ID,
		FolderName: newAlbum.FolderName,
		CreatedAt:  newAlbum.CreatedAt,
	})
}

func (h *AlbumHandler) GetAlbum(ctx *fiber.Ctx) error {
	queryParam := util.GetParams(ctx, "album_id")
	if queryParam == "" {
		log.Println("no params provided")
		return ctx.JSON("no params provided")
	}

	albumID, err := strconv.Atoi(queryParam)
	if err != nil {
		return ctx.JSON(err)
	}

	album, err := h.albumService.GetAlbum(albumID)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(&dto.AlbumResponse{
		FolderName: album.FolderName,
		CreatedAt:  album.CreatedAt,
	})
}
