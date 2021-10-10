package util

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetRequestUserID(ctx *fiber.Ctx) (int, error) {
	return strconv.Atoi(ctx.Get("user", ""))
}

func GetParams(ctx *fiber.Ctx, params string) string {
	return ctx.Params(params, "")
}

func GetQueryParams(ctx *fiber.Ctx, params string) string {
	return ctx.Query(params, "")
}
