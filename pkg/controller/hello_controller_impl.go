package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type helloControllerImpl struct{}

func (helloControllerImpl) Hello(w http.ResponseWriter, r *http.Request) {
	msg := []byte(`{"message": "hello"}`)
	w.Header().Add("Content-Length", strconv.Itoa(len(msg)))
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	if _, err := w.Write(msg); err != nil {
		fmt.Fprintf(os.Stderr, "{\"error\":%q}\n", err)
	}
}
