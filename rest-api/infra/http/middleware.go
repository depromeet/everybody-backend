package http

import (
	"encoding/json"
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorType string `json:"error_type"`
}

// errorHandle 은 최종적으로 리턴된 error들에 대해 적절한 응답을 생성하고
// 예상치 못한 에러의 경우 500 에러로 응답합니다.
func errorHandle(ctx *fiber.Ctx, err error) error {
	rootErr := errors.GetRootStackError(err)
	log.Errorf("%+v", rootErr)

	notFoundErr := new(service.NotFoundError)
	unmarshalTypeErr := new(json.UnmarshalTypeError)

	if errors.As(err, notFoundErr) {
		return ctx.Status(404).JSON(newErrorResponse("리소스를 찾을 수 없습니다.", reflect.TypeOf(err).Name()))
	} else if errors.As(err, &unmarshalTypeErr) {
		e := errors.Cause(err).(*json.UnmarshalTypeError) // .Cause()는 .Cause()를 구현하지 않은 에러를 error로 리턴합니다.
		return ctx.Status(400).JSON(newErrorResponse(fmt.Sprintf("%s에 대한 잘못된 타입의 값입니다. %s 타입을 이용해주세요.", e.Field, e.Type.Name()), reflect.TypeOf(e).Name()))
	} else {
		return ctx.Status(500).JSON(newErrorResponse("알 수 없는 에러가 발생했습니다. 에브리바디에 문의해주세요.", "internalError"))
	}
}

func defaultLog(ctx *fiber.Ctx) error {
	log.Infof("Request Headers: %s\n", ctx.Request().Header.String())
	if strings.HasPrefix(ctx.Get("Content-Type", ""), "application/json") {
		log.Infof("Request JSON Body: %s\n", ctx.Body())
	}
	return ctx.Next()
}

func newErrorResponse(message, errorType string) *ErrorResponse {
	return &ErrorResponse{
		Message:   message,
		ErrorType: errorType,
	}
}
