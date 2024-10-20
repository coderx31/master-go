package main

import (
	"io"
	"net/http"
)

/**
test pyramid
	E2E tests
	Integration tests
	Unit tests
*/

/*
*
- most common way to classify the tests os using built tags.
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}
func main() {

}
