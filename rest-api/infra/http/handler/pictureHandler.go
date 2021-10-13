package handler

import (
	"errors"
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
	// var pictureReq dto.PictureRequest
	// err := ctx.BodyParser(&pictureReq)
	// if err != nil {
	// 	return err
	// }
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return err
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	pictureReq := dto.PictureMultiPart{
		AlbumID:  form.Value["album_id"],
		BodyPart: form.Value["body_part"],
		File:     form.File["picture"],
	}

	picture, err := h.pictureService.SavePicture(userID, &pictureReq)
	if err != nil {
		return err
	}

	return ctx.JSON(picture)
}

// GetAllPictures는 user가 가지고 있는 모든 사진 조회
func (h *PictureHandler) GetAllPictures(ctx *fiber.Ctx) error {
	// TODO: albumid가 아니라 userid로 전체 사진 조회하도록 수정
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return err
	}

	pictures, err := h.pictureService.GetAllPictures(userID)
	if err != nil {
		return err
	}

	return ctx.JSON(pictures)
}

func (h *PictureHandler) GetPicture(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "picture_id")
	if param == "" {
		return errors.New("params should be provided")
	}

	pictureID, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	picture, err := h.pictureService.GetPicture(pictureID)
	if err != nil {
		return err
	}

	return ctx.JSON(picture)
}
