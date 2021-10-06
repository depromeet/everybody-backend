package util

import "github.com/gofiber/fiber/v2"

func GetRequestUserID(ctx *fiber.Ctx) string {
	return ctx.Get("user", "")
}

func GetParams(ctx *fiber.Ctx, params string) string {
	return ctx.Params(params, "")
}

func GetQueryParams(ctx *fiber.Ctx, params string) string {
	return ctx.Query("params", "")
}
