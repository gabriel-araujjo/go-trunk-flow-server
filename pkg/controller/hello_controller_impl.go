package controller

import (
	"net/http"
	"strconv"
)

type helloControllerImpl struct{}

func (helloControllerImpl) Hello(w http.ResponseWriter, r *http.Request) {
	msg := []byte(`{"message": "hello"}`)
	w.Header().Add("Content-Length", strconv.Itoa(len(msg)))
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(msg)
}
