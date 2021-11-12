package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
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

// 몇몇 타입들의 에러들인 wrong value: negative value: blah blah
// 이런 식으로 colon을 이용해서 에러 메시지를 재귀적으로 표현함 ㅜㅜ
// 클라가 에러 메시지를 유저에게 바로 띄워주려면 이쁜 메시지를 전달해줘야함..
func GetErrorMessageForClient(message string) string {
	splitted := strings.Split(message, ":")
	trimmed := strings.TrimSpace(splitted[0])
	if len(trimmed) == 0 {
		return GetErrorMessageForClient(strings.Join(splitted[1:], ":"))
	}

	return trimmed
}
