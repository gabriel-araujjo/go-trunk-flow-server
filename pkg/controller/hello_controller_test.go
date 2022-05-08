package controller

import (
	"bytes"
	"mime"
	"net/http/httptest"
	"testing"
)

func TestHelloController(t *testing.T) {
	r := httptest.NewRequest("GET", "/hello", bytes.NewBufferString(""))
	w := httptest.NewRecorder()

	hello := NewHelloController()
	hello.Hello(w, r)

	result := w.Result()

	if result.StatusCode != 200 {
		t.Errorf("expecting status 200, but got %d", result.StatusCode)
	}

	mediaType, _, err := mime.ParseMediaType(result.Header.Get("Content-Type"))

	if err != nil {
		t.Errorf("error parsing mimetype: %v", err)
	}

	if mediaType != "application/json" {
		t.Errorf("expecting a json, but got %q", mediaType)
	}

	if result.ContentLength <= 0 {
		t.Errorf("expecting content-length is set")
	}
}
