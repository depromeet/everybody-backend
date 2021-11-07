package service

import (
	"github.com/depromeet/everybody-backend/rest-api/ent"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NotFoundError(t *testing.T) {
	var entErr *ent.NotFoundError
	notFoundErrorPtr := new(NotFoundError)
	notFoundError := NotFoundError(entErr)
	assert.ErrorAs(t, &ent.NotFoundError{}, notFoundErrorPtr)
	assert.ErrorAs(t, &ent.NotFoundError{}, &notFoundError)
	assert.Panics(t, func() {
		// panic: errors: *target must be interface or implement error
		// 타겟은 포인터형이어야한다.
		errors.As(&ent.NotFoundError{}, notFoundError)
	})
}
