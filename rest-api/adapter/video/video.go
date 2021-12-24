package video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
)

type downloadVideoRequest struct {
	Keys     []imageKey `json:"keys"`
	Duration *float64   `json:"duration"`
}

type imageKey struct {
	Key string `json:"key"`
}

// 네이밍을 어떻게 하는 게 좋을까..
// hexagonal architecture에 따라 Port라고 명명해봤음.
// unittest할 때 mocking하기 위해 interface로 선언
type VideoPort interface {
	DownloadVideo(user int, imageKeys []string, duration *float64) (io.Reader, error)
}

type videoAdapter struct{}

func NewVideoPort() VideoPort {
	return &videoAdapter{}
}

func (a *videoAdapter) DownloadVideo(user int, imageKeys []string, duration *float64) (io.Reader, error) {
	if duration == nil {
		tmp := 0.25
		duration = &tmp
	}
	data := downloadVideoRequest{
		Keys:     []imageKey{},
		Duration: duration,
	}
	for _, key := range imageKeys {
		data.Keys = append(data.Keys, imageKey{Key: key})
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/video", config.Config.Video.APIRoot), bytes.NewReader(body))
	req.Header.Add("user", strconv.Itoa(user))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	c := new(http.Client)
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != 200 {
		respBuf := []byte{}
		_, err = resp.Body.Read(respBuf)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, errors.New(string(respBuf))
	}

	return resp.Body, nil
}
