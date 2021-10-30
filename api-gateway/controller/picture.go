package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"

	"github.com/depromeet/everybody-backend/api-gateway/config"
	"github.com/depromeet/everybody-backend/api-gateway/util"
)

type PictureController struct {
}

func UploadPicture(c echo.Context) error {
	// 람다 호출에 앞서 인증절차 선행 - request를 새로 만들 것이므로 기존 헤더는 신경쓰지 않아도됨
	userId := 0
	if config.Config.ApiGw.AuthEnable {
		token := c.Request().Header.Get("Authorization")
		id, err := util.VerifyAccessToken(token)
		userId = id
		if err != nil {
			return c.String(http.StatusForbidden, "Token invalid")
		}
	} else {
		log.Warn("Auth DISABLED... func 'UploadPicture' processing with userId=0")
	}

	log.Info("UploadPicture -> userId=", userId)

	// 람다 호출해서 이미지 업로드
	// TODO: multipart 이미지 획득해서 req 만드는 부분 구현 필요
	lambdaResp, err := http.Get("http://localhost:5000/health") // TODO: config에 있는 URL,method로 userId 넣어서 호출하는걸로 변경
	if err != nil {
		log.Error("image upload fail...", err)
		return c.String(http.StatusInternalServerError, "image upload fail...")
	}
	defer lambdaResp.Body.Close()
	data, err := ioutil.ReadAll(lambdaResp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	lambdaResBody := string(data)

	// 람다 호출 결과 확인
	if lambdaResp.StatusCode == 200 {
		// 성공한 경우 새로운 request를 만들어서 rest-api 호출
		// TODO: 이부분 필요한 데이터만 파싱해서 나가도록 수정
		values := map[string]string{"body_part": c.FormValue("body_part"), "album_id": c.FormValue("album_id")} // TODO: 이거 제대로 동작하는지 확인
		jsonValue, _ := json.Marshal(values)

		panic("aaaaaa")

		req, _ := http.NewRequest("", "", bytes.NewBuffer(jsonValue)) // method and url will be set bottom
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("user", strconv.Itoa(userId))

		code, resBody := callRestApi(c, false, req, "/pictures", "POST") // rest-api '서버의 업로드한 사진 데이터 저장' api path
		if code == http.StatusOK {
			return c.String(code, resBody)
		} else {
			// 람다 업로드는 성공했으나 rest-api 호출이 실패한 경우...
			// TODO: fatal로 프로세스 죽기 전에 S3에서 이미 업로드된 사진을 삭제해줘야함
			log.Fatal("rest-api Request fail... code=" + strconv.Itoa(code) + " resBody=" + resBody)
			return c.String(code, "rest-api Request fail...")
		}

	} else {
		// 람다 호출 실패한 경우에 처리
		log.Fatal("image upload fail... code=" + strconv.Itoa(lambdaResp.StatusCode) + " body=" + lambdaResBody)
		return c.String(http.StatusInternalServerError, "image upload fail...")
	}
}
