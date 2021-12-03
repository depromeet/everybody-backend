package util

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// 참고: https://pkg.go.dev/encoding/json#NewDecoder
// fiber의 application/json에 대한 default decoder는 json unmarshaller
type CustomTime time.Time

// ct에 저장
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return errors.WithStack(err)
	}
	if t, err := time.Parse("2006-01-02 15:04:05", s); err != nil {
		return errors.WithStack(err)
	} else {
		// 전달받은 건 한국시인데 해석할 때에는 같은 절대 시간값으로 UTC로 해석하기 때문에 9시간을 더해준다.
		t.Add(9 * time.Hour)
		*ct = CustomTime(t)
		return nil
	}
}

// ct를 []byte에 저장
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct)
}

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

func ConvertTimeToStr(t time.Time) (year, month, day, hour, min int) {
	return t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute()
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
