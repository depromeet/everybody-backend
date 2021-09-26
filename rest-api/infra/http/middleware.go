package http

import (
    "errors"
    "github.com/depromeet/everybody-backend/rest-api/ent"
    "github.com/gofiber/fiber/v2"
    "reflect"
)


type ErrorResponse struct{
    Message string `json:"message"`
    ErrorType string `json:"error_type"`
}

// errorHandle 은 최종적으로 리턴된 error들에 대해 적절한 응답을 생성하고
// 예상치 못한 에러의 경우 500 에러로 응답합니다.
func errorHandle(ctx *fiber.Ctx, err error) error{
    if err != nil {
        notFoundErr := &ent.NotFoundError{}
        if errors.As(err, &notFoundErr){
            return ctx.Status(404).JSON(e("리소스를 찾을 수 없습니다.", reflect.TypeOf(err).Name()))
        }
    }

    return ctx.Status(500).JSON(e("알 수 없는 에러가 발생했습니다. 에브리바디에 문의해주세요.", "internalError"))
}

// e 는 ErrorResponse를 간단히 생성하기 위한 Shortcut
func e(message, errorType string) *ErrorResponse{
    return &ErrorResponse{
        Message: message,
        ErrorType: errorType,
    }
}