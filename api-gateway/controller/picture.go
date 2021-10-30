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
	// 람다 호출에 앞서 인증절차 선행 - request를 새로 만들 것이므로 헤더는 신경쓰지 않아도됨
	userId := uint64(0)
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
	lambdaResp, err := http.Get("http://localhost:5000/health") // TODO: config에 있는 URL,method로 userId 넣어서 호출하는걸로 변경
	if err != nil {
		log.Error("image upload fail...", err)
		return c.String(http.StatusInternalServerError, "image upload fail...")
	}
	data, err := ioutil.ReadAll(lambdaResp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	lambdaResBody := string(data)

	// 람다 호출 결과 확인
	if lambdaResp.StatusCode == 200 {
		// 성공한 경우 새로운 request를 만들어서 rest-api 호출
		values := map[string]string{"body_part": c.FormValue("body_part"), "album_id": c.FormValue("album_id")} // TODO: 이거 제대로 동작하는지 확인
		jsonValue, _ := json.Marshal(values)
		req, _ := http.NewRequest("", "", bytes.NewBuffer(jsonValue)) // method and url will be set bottom
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("user", strconv.FormatUint(userId, 10)) // to decimal string

		code, s := callRestApi(c, false, req, "/pictures", "POST") // rest-api '서버의 업로드한 사진 데이터 저장' api path
		if code == http.StatusOK {
			return c.String(code, s)
		} else {
			// 람다 업로드는 성공했으나 rest-api 호출이 실패한 경우...
			log.Error("rest-api Request fail... code=" + strconv.Itoa(code) + " body=" + s)
			return c.String(code, "image upload fail...")
		}

	} else {
		// 람다 호출 실패한 경우에 처리
		log.Error("image upload fail... code=" + strconv.Itoa(lambdaResp.StatusCode) + " body=" + lambdaResBody)
		return c.String(http.StatusInternalServerError, "image upload fail...")
	}
}
