package main

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	s := Server{}
	s.Hello(w, req)

	if w.Result().StatusCode != 200 {
		t.Fatalf("Unexpected status code %d", w.Result().StatusCode)
	}

	body := w.Body.String()
	if body != `{"message": "Hello World"}` {
		t.Fatalf("Unexpected body received: %s", body)
	}
}
