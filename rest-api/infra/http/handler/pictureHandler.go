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
	albumService   service.AlbumServiceInterface
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

	pictureReq := new(dto.CreatePictureRequest)
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
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	param := util.GetParams(ctx, "picture_id")
	if len(param) == 0 {
		return errors.WithStack(errors.New("picture_id params를 입력해주세요"))
	}

	pictureID, err := strconv.Atoi(param)
	if err != nil {
		return errors.WithStack(err)
	}

	picture, err := h.pictureService.GetPicture(userID, pictureID)
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

	pictureReq := new(dto.GetPictureRequest)
	err = ctx.QueryParser(pictureReq)
	if err != nil {
		return errors.WithStack(err)
	}

	// query string으로 uploader, album이 둘 다 없다면 잘못된 요청
	if len(pictureReq.Uploader) == 0 && len(pictureReq.Album) == 0 {
		return &BadRequestError{errors.New("적절한 query string으로 요청해주세요")}
	}

	pictures, err := h.pictureService.GetAllPictures(userID, pictureReq)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(pictures)
}

func (h *PictureHandler) DeletePicture(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.Wrap(err, "잘못된 유저 ID입니다.")
	}

	pictureID, err := strconv.Atoi(ctx.Params("picture_id"))
	if err != nil {
		return errors.Wrap(err, "올바르지 않은 사진 ID입니다.")
	}

	err = h.pictureService.Delete(userID, pictureID)
	if err != nil {
		return errors.WithMessage(err, "")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
