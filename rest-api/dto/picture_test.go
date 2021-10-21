package dto

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_createImageURL(t *testing.T) {
	// 미리 버킷에 업로드해놓은 sample.jpg를 이용해 테스트 해봄
	url, err := createImageURL("sample.jpg")
	assert.NoError(t, err)
	t.Log(url)
	resp, err := http.Get(url)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
