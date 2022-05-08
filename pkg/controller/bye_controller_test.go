//go:build controller.bye || dev

package controller

import (
	"bytes"
	"mime"
	"net/http/httptest"
	"testing"
)

func TestByeController(t *testing.T) {
	r := httptest.NewRequest("GET", "/bye", bytes.NewBufferString(""))
	w := httptest.NewRecorder()

	bye := NewByeController()
	bye.Bye(w, r)

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
}
