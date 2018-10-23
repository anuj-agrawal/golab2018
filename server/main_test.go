package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandler(b *testing.B) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/names?id=abc", nil)
	for i := 0; i < b.N; i++ {
		handler(w, r)
	}
}
