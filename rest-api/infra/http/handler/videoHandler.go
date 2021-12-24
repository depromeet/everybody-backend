package handler

import (
	"strconv"

	"github.com/depromeet/everybody-backend/rest-api/dto"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/depromeet/everybody-backend/rest-api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type VideoHandler struct {
	videoService service.VideoServiceInterface
}

func NewVideoHandler(videoService service.VideoServiceInterface) *VideoHandler {
	return &VideoHandler{
		videoService: videoService,
	}
}

func (h *VideoHandler) DownloadVideo(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	body := new(dto.DownloadVideoRequest)
	err = ctx.BodyParser(body)
	if err != nil {
		return errors.WithStack(err)
	}

	video, err := h.videoService.DownloadVideo(userID, body)
	if err != nil {
		return err
	}
	ctx.Response().Header.Add("Content-Type", "video/mp4")

	return ctx.SendStream(video)
}

func (h *VideoHandler) SaveVideo(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	videoReq := new(dto.VideoRequest)
	err = ctx.BodyParser(videoReq)
	if err != nil {
		return errors.WithStack(err)
	}

	video, err := h.videoService.SaveVideo(userID, videoReq)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(video)
}

func (h *VideoHandler) GetVideo(ctx *fiber.Ctx) error {
	params := util.GetParams(ctx, "video_id")
	if params == "" {
		return errors.New("video_id params should be provided")
	}

	videoID, err := strconv.Atoi(params)
	if err != nil {
		return errors.WithStack(err)
	}

	video, err := h.videoService.GetVideo(videoID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(video)
}

func (h *VideoHandler) GetAllVideos(ctx *fiber.Ctx) error {
	userID, err := util.GetRequestUserID(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	videos, err := h.videoService.GetAllVideos(userID)
	if err != nil {
		return errors.WithStack(err)
	}

	return ctx.JSON(videos)
}
