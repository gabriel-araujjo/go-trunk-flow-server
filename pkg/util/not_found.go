package util

import (
	"net/http"
	"strconv"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	msg := []byte(`{"error": "not found"}`)

	w.Header().Add("Content-Length", strconv.Itoa(len(msg)))
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(404)
	w.Write(msg)
}
