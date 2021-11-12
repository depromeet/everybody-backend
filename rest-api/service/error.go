package service

import (
	goerrors "errors"
)

type NotFoundError error

var (
	ForbiddenError = goerrors.New("권한이 없습니다.")
	// 왜 이걸 Unauthenticated로 안하고 Unauthorized라고 하는걸까..
	UnauthorizedError = goerrors.New("인증되지 않은 사용자입니다.")
)
