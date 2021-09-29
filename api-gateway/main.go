package main

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":8443", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Info("called! -> " + req.URL.Path)

	// GET 호출
	resp, err := http.Get("http://localhost:8888/testapi") // for test...
	if err != nil {
		log.Error(err)
		panic(err)
	}

	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	str := string(data)
	log.Debug(str)

	w.Write([]byte(str))
}
