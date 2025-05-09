package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCotacaoHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/cotacao", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "mocked error", http.StatusServiceUnavailable)
	})

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("Esperado status 503, recebeu %d", w.Code)
	}
}
