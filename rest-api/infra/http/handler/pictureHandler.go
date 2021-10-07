package handler

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
)

type PictureHandler struct {
	pictureService service.PictureServiceInterface
}

func NewPictureHandler(pictureService service.PictureServiceInterface) *PictureHandler {
	return &PictureHandler{
		pictureService: pictureService,
	}
}

func (h *PictureHandler) SavePicture(ctx *fiber.Ctx) error {
	var pictureReq dto.PictureRequest
	err := ctx.BodyParser(&pictureReq)
	if err != nil {
		return ctx.JSON(err)
	}

	result, err := h.pictureService.SavePicture(&pictureReq)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(result)
}

func (h *PictureHandler) GetAllPictures(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "album_id")
	if param == "" {
		ctx.JSON("params should be provided")
	}

	albumID, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(err)
	}

	pictures, err := h.pictureService.GetAllPicutres(albumID)
	if err != nil {
		return ctx.JSON(err)
	}

	picturesResponse := make(dto.PicturesResponse, 0)
	for _, picture := range pictures {
		var pictureResponse *dto.PictureResponse

		pictureResponse.ID = picture.ID
		pictureResponse.CreatedAt = picture.CreatedAt
		// s3 url 이나 hashed file name 등등 picture의 메타 데이터 주어야 함.

		picturesResponse = append(picturesResponse, *pictureResponse)
	}

	return ctx.JSON(picturesResponse)
}

func (h *PictureHandler) GetPicture(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "picture_id")
	if param == "" {
		ctx.JSON("params should be provided")
	}

	pictureID, err := strconv.Atoi(param)
	if err != nil {
		return ctx.JSON(err)
	}

	picture, err := h.pictureService.GetPicture(pictureID)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(dto.PictureResponse{
		CreatedAt: picture.CreatedAt,
		// picture의 메타 데이터 주어야 함.
	})
}
