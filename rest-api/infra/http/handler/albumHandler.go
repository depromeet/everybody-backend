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

func (h *AlbumHandler) GetAllAlbums(ctx *fiber.Ctx) error {
	userID := util.GetRequestUserID(ctx)

	albums, err := h.albumService.GetAllAlbums(userID)
	if err != nil {
		return ctx.JSON(err)
	}

	albumsResponse := make(dto.AlubumsResponse, 0)
	for _, album := range albums {
		var albumResponse dto.AlbumResponse
		albumResponse.ID = album.ID
		albumResponse.FolderName = album.FolderName
		albumResponse.CreatedAt = album.CreatedAt

		albumsResponse = append(albumsResponse, albumResponse)
	}

	return ctx.JSON(albumsResponse)
}

func (h *AlbumHandler) GetAlbum(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "album_id")
	if param == "" {
		log.Println("no params provided")
		return ctx.JSON("no params provided")
	}

	albumID, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(err)
	}

	// GetAlbum 에서 각 앨범에 해당하는 pictures도 조회해야 함
	album, err := h.albumService.GetAlbum(albumID)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(&dto.AlbumResponse{
		FolderName: album.FolderName,
		// PictureList: []picture,
		CreatedAt: album.CreatedAt,
	})
}
