package handler

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
	if len(param) == 0 {
		return errors.WithStack(errors.New("picture_id params를 입력해주세요"))
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

// GetAllPictures는 query string에 따른 사진들 조회(uploader=? / album=?&body_part=?))
func (h *PictureHandler) GetAllPictures(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	pictureQueryString := new(dto.PictureQueryString)
	err = ctx.QueryParser(pictureQueryString)
	if err != nil {
		return errors.WithStack(err)
	}

	// query string으로 uploader가 없다면 잘못된 요청
	if len(pictureQueryString.Uploader) == 0 {
		return errors.WithStack(errors.New("적절한 query string으로 요청해주세요"))
	}

	pictures, err := h.pictureService.GetAllPictures(userID, pictureQueryString)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(pictures)
}
