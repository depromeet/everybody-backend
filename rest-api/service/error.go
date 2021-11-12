package service

import "github.com/pkg/errors"

type NotFoundError error

var (
	ForbiddenError = errors.New("권한이 없습니다.")
	// 왜 이걸 Unauthenticated로 안하고 Unauthorized라고 하는걸까..
	UnauthorizedError = errors.New("인증되지 않은 사용자입니다.")
)
