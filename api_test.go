package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks__API(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books []Book
	json.Unmarshal([]byte(w.Body.String()), &books)

	assert.Equal(t, books, getBooks())
}

func TestGetBook_ByName_OK__API(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/H2G2?crit=name", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books []Book
	json.Unmarshal([]byte(w.Body.String()), &books)

	assert.Equal(t, books[0].Name, "H2G2")
	assert.Equal(t, books[0].Isbn, "h2g242")
}

func TestGetBook_ByName_nocrit_OK__API(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/H2G2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books []Book
	json.Unmarshal([]byte(w.Body.String()), &books)

	assert.Equal(t, len(books), 0)
}

func TestGetBookBy_Isbn_OK__API(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/lotr42?crit=isbn", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books []Book
	json.Unmarshal([]byte(w.Body.String()), &books)

	assert.Equal(t, books[0].Name, "Lord of the rings")
	assert.Equal(t, books[0].Isbn, "lotr42")
}

func TestGetBookBy_Isbn_nocrit_OK__API(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/lotr42", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var books []Book
	json.Unmarshal([]byte(w.Body.String()), &books)

	assert.Equal(t, books[0].Name, "Lord of the rings")
	assert.Equal(t, books[0].Isbn, "lotr42")
}
