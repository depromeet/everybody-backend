package handler

import (
	"strconv"

	"github.com/pkg/errors"

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
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	pictureReq := new(dto.PictureRequest)
	err = ctx.BodyParser(pictureReq)
	if err != nil {
		return errors.WithStack(err)
	}

	picture, err := h.pictureService.SavePicture(userID, pictureReq)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(picture)
}

func (h *PictureHandler) GetPicture(ctx *fiber.Ctx) error {
	param := util.GetParams(ctx, "picture_id")
	if param == "" {
		return errors.New("picture_id params should be provided")
	}

	pictureID, err := strconv.Atoi(param)
	if err != nil {
		return errors.WithStack(err)
	}

	picture, err := h.pictureService.GetPicture(pictureID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(picture)
}

// GetAllPictures는 user가 가지고 있는 모든 사진 조회
func (h *PictureHandler) GetAllPictures(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	pictures, err := h.pictureService.GetAllPictures(userID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(pictures)
}

// GetPictures는 특정 앨범의 특정 신체 부위의 사진 조회(영상 만들때도 필요)
func (h *PictureHandler) GetPictures(ctx *fiber.Ctx) error {
	albumQuery := util.GetQueryParams(ctx, "album_id")
	if albumQuery == "" {
		return errors.New("album_id query param should be provided")
	}

	albumID, err := strconv.Atoi(albumQuery)
	if err != nil {
		return errors.WithStack(err)
	}

	bodyPartQuery := util.GetQueryParams(ctx, "body_part")

	pictures, err := h.pictureService.GetPictures(albumID, bodyPartQuery)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(pictures)
}
