package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	expected := "id: 1, param: 2"
	req := httptest.NewRequest("GET", "/1?param=2", nil)
	w := httptest.NewRecorder()

	rootHandler(w, req)
	resp := w.Result()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error reading response")
	}
	responseStr := string(bytes)
	t.Log(responseStr)
	// todo why id is empty?
	if expected != responseStr {
		t.Fatal("response doesn't match expected")
	}
}
