package service

import (
	"github.com/pkg/errors"
	"testing"
)

func Test_NotFoundError(t *testing.T) {
	err := err3()
	//notFoundErr := NotFoundError(err)
	//t.Logf("%+v", notFoundErr)
	t.Logf("%+v", err)

}

func err1() error {
	return errors.New("err1")
}

func err2() error {
	return errors.WithStack(err1(), "err2")
}

func err3() error {
	return errors.WithStack(err2(), "err3")
}
