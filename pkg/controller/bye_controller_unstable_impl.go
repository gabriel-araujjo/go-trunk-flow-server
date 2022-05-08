//go:build controller.bye || dev

package controller

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gabriel-araujjo/go-trunk-flow-server/pkg/features"
)

type byeControllerImpl struct{}

func (byeControllerImpl) Bye(w http.ResponseWriter, r *http.Request) {
	var msg []byte
	var status int

	if rand.Intn(2) == 0 {
		msg = []byte(`{"error": "internal server error"}`)
		status = 500
	} else {
		msg = []byte(`{"message": "bye"}`)
		status = 200
	}

	w.Header().Add("Content-Length", strconv.Itoa(len(msg)))
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(msg)
}

func init() {
	features.Features = append(features.Features, "controller.bye")
}
