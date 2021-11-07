package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
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

func ConvertIntToTime(year, month, day int) (time.Time, error) {
	return time.Parse(time.RFC3339, fmt.Sprintf("%04d-%02d-%02dT12:00:00+09:00", year, month, day))
}

func ConvertTimeToStr(t time.Time) (year, month, day int) {
	return t.Year(), int(t.Month()), t.Day()
}
