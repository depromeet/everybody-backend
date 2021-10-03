package util

import "github.com/gofiber/fiber/v2"

func GetRequestUserID(ctx *fiber.Ctx) string {
	return ctx.Get("user", "")
}
