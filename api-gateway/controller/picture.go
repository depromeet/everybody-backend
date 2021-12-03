package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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

	// 필요한 파라미터 값들 획득
	albumId, err := strconv.Atoi(c.FormValue("album_id"))
	if err != nil || albumId < 1 {
		log.Error("albumId parse error... albumId=", c.FormValue("album_id"), " err=", err)
		return c.String(http.StatusBadRequest, "albumId parse error...")
	}

	bodyPart := c.FormValue("body_part")
	if bodyPart == "" {
		log.Error("bodyPart parse error... bodyPart=", c.FormValue("body_part"))
		return c.String(http.StatusBadRequest, "bodyPart parse error...")
	}

	imageFile, err := c.FormFile("image")
	if err != nil {
		log.Error("get Formfile fail... err=", err)
		return c.String(http.StatusBadRequest, "get Formfile fail...")
	}

	takenAt := c.FormValue("taken_at")
	if err != nil {
		log.Error("takenAt parse error... takenAt=", c.FormValue("takenAt"))
		return c.String(http.StatusBadRequest, "takenAt parse error...")
	}

	// 람다 호출해서 이미지 업로드
	lambdaResCode, _, lambdaResBody := callLambdaImageUpload(userId, imageFile)
	if lambdaResCode != 200 {
		// 람다 업로드 실패시 프로세스 죽음에 주의(fatal)
		log.Fatal("image upload fail... code=" + strconv.Itoa(lambdaResCode) + " err=" + lambdaResBody)
		// TODO: 에러 처리방안 논의 후 fatal -> panic? error? 로 수정 필요
		return c.String(http.StatusInternalServerError, "image upload fail... err="+lambdaResBody)

	} else {
		log.Info("image upload success! lambdaResBody=", lambdaResBody)
	}

	// 성공한 람다 호출 결과 파싱
	type dataType struct {
		Keys []string `json:"keys"`
	}
	type lambdaResBodyType struct {
		Data  dataType `json:"data"`
		Error string   `json:"error"` // "error":null 이면 이 필드는 ""로 변환됨
	}
	var lambdaResObj lambdaResBodyType
	json.Unmarshal([]byte(lambdaResBody), &lambdaResObj)

	// 람다 호출 후 받은 res에서 key를 획득하여 새롭게 rest-api 호출요청
	reqMap := make(map[string]interface{})
	reqMap["key"] = lambdaResObj.Data.Keys[0] // 현재는 사진 1장씩만 해서 리스트의 첫번째 원소만 업로드함...
	reqMap["album_id"] = albumId
	reqMap["body_part"] = bodyPart
	reqMap["taken_at"] = takenAt
	reqMapByte, _ := json.Marshal(reqMap)

	req, _ := http.NewRequest("", "", bytes.NewBuffer(reqMapByte)) // method and url will be set bottom
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user", strconv.Itoa(userId))
	code, header, resBody := callRestApi(c, req, "/pictures", "POST")

	if code != http.StatusOK {
		// TODO: 200 아니면 S3에서 이미 업로드된 사진을 삭제해줘야함

		log.Panic("rest-api Request fail... code=" + strconv.Itoa(code) + " resBody=" + resBody)
		// TODO: middleware recover 도입.. https://echo.labstack.com/middleware/recover/
		// TODO: 에러 처리방안 논의 후 수정 필요.. 일단 그냥 if 바깥으로 빠져서 rest-api한테 받은거 그대로 리턴하게함 //return c.JSON(http.StatusInternalServerError, resBody)
	}

	// rest에게 받은 응답을 그대로 전달
	if header != nil {
		for k, v := range header {
			c.Response().Header().Set(k, v.(string))
		}
	}
	return c.String(code, resBody)
}

func callLambdaImageUpload(userId int, imageFile *multipart.FileHeader) (int, map[string]interface{}, string) {
	// ref: https://stackoverflow.com/questions/44302374/can-i-post-with-content-type-multipart-form-data
	// ref: https://github.com/golang/go/issues/30218

	// body buffer 생성
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// Text형태의 key-value 추가 - 지금은 쓰고있는 곳 없음...
	//bodyWriter.WriteField("mykey", "myval")

	// File형태의 key-value 추가
	fileWriter, err := bodyWriter.CreatePart(imageFile.Header)
	if err != nil {
		log.Error("bodyWriter.CreatePart fail... imageFile.Header=", imageFile.Header, " err=", err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	fileContent, err := imageFile.Open()
	if err != nil {
		log.Error("imageFile.Open fail... imageFile=", imageFile, " err=", err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	byteContainer, err := ioutil.ReadAll(fileContent)
	if err != nil {
		log.Error("ioutil.ReadAll fail... fileContent=", fileContent, " err=", err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	fileWriter.Write(byteContainer)

	bodyWriter.Close()

	// request 생성 및 헤더 설정
	req, _ := http.NewRequest(config.Config.TargetServer.LambdaImageUpload.Method, config.Config.TargetServer.LambdaImageUpload.Address, bodyBuf)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType()) // "multipart/form-data; boundary=170b98f8b82b....."
	req.Header.Set("user", strconv.Itoa(userId))

	// 실제 호출
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req) // TODO: util/httpclient로 대체 필요... 커넥션 풀 및 타임아웃 제어 필요..
	if err != nil {
		log.Error("call lambda fail... err=", err)
		return http.StatusInternalServerError, nil, err.Error()
	}
	defer resp.Body.Close()

	// 응답 받은 결과 복사 - Header
	h := make(map[string]interface{})
	for k := range resp.Header {
		h[k] = resp.Header.Get(k)
	}

	// 응답 받은 결과 복사 - Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp.Body read error... err=", err)
		panic(err) // TODO: log.panic으로 바꾸고.. middleware recover 도입.. https://echo.labstack.com/middleware/recover/
	}

	return resp.StatusCode, h, string(data)
}
