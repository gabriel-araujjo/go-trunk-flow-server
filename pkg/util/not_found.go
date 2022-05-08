package util

import (
	"fmt"
	"net/http"
	"os"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	msg := []byte(`{"error": "not found"}`)

	w.WriteHeader(404)
	if _, err := w.Write(msg); err != nil {
		fmt.Fprintf(os.Stderr, "{\"error\":%q}\n", err)
	}
}
