package controller

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type VideoController struct{}

func DownloadVideo(c echo.Context) error {
	userId := 0
	if config.Config.ApiGw.AuthEnable {
		token := c.Request().Header.Get("Authorization")
		id, err := util.VerifyAccessToken(token)
		userId = id
		if err != nil {
			return c.String(http.StatusForbidden, "Token invalid")
		}
	} else {
		log.Warn("Auth DISABLED... func 'DownloadVideo' processing with userId=0")
	}

	log.Info("DownloadVideo -> userId=", userId)

	// 람다 호출해서 이미지 업로드
	lambdaResCode, h, lambdaResBody, err := callLambdaDownloadVideo(userId, c.Request().Body)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if lambdaResCode != http.StatusOK {
		msg, _ := ioutil.ReadAll(lambdaResBody)
		log.Error("download image failed... code=" + strconv.Itoa(lambdaResCode) + " resp=" + string(msg))
		return c.String(http.StatusInternalServerError, string(msg))
	}

	// rest에게 받은 응답을 그대로 전달
	if h != nil {
		for k, v := range h {
			c.Response().Header().Set(k, v.(string))
		}
	}

	return c.Stream(lambdaResCode, "video/mp4", lambdaResBody)
}

func callLambdaDownloadVideo(userId int, reader io.Reader) (int, map[string]interface{}, io.Reader, error) {
	// request 생성 및 헤더 설정
	req, _ := http.NewRequest(config.Config.TargetServer.LambdaVideoDownload.Method, config.Config.TargetServer.LambdaVideoDownload.Address, reader)
	req.Header.Set("Content-Type", "application/json") // "multipart/form-data; boundary=170b98f8b82b....."
	req.Header.Set("user", strconv.Itoa(userId))

	// 실제 호출
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error("call lambda fail... err=", err)
		return http.StatusInternalServerError, nil, nil, err
	}
	defer resp.Body.Close()

	tmp, err := ioutil.ReadAll(resp.Body)
	respBody := bytes.NewReader(tmp)

	// 응답 받은 결과 복사 - Header
	h := make(map[string]interface{})
	for k := range resp.Header {
		h[k] = resp.Header.Get(k)
	}

	return resp.StatusCode, h, respBody, nil
}
