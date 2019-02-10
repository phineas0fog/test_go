package main

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGemEM(t *testing.T) {
	msg := getEM(404)
	if msg != "This book does not exists" {
		t.Errorf("Error message is incorrect")
	}
}

func TestGemEM_UNKN(t *testing.T) {
	msg := getEM(0)
	if msg != "UNKNOWN" {
		t.Errorf("Error message is incorrect")
	}
}

func TestSetupRouter(t *testing.T) {
	router := setupRouter()
	sample := gin.New()

	routerT := reflect.TypeOf(router).Kind()
	sampleT := reflect.TypeOf(sample).Kind()

	if routerT != sampleT {
		t.Errorf("Return type is incorrect")
	}

}

func TestGetBooks(t *testing.T) {
	books := getBooks()

	if len(books) != 2 {
		t.Errorf("function must return 2 books")
	}
}

func TestGetBookBy_Name_OK(t *testing.T) {
	books, err := getBook("name", "H2G2")

	if err != nil {
		t.Errorf("Function returned an error, got %s", err)
	}

	book := books[0]

	if book.Name != "H2G2" {
		t.Errorf("Returned Name is incorrect, got %s", book.Name)
	}

	if book.Isbn != "h2g242" {
		t.Errorf("Returned Isbn is incorrect, got %s", book.Isbn)
	}
}

func TestGetBookBy_Isbn_OK(t *testing.T) {
	books, err := getBook("isbn", "lotr42")

	if err != nil {
		t.Errorf("Function returned an error, got %s", err)
	}

	book := books[0]

	if book.Name != "Lord of the rings" {
		t.Errorf("Returned Name is incorrect, got %s", book.Name)
	}

	if book.Isbn != "lotr42" {
		t.Errorf("Returned Isbn is incorrect, got %s", book.Isbn)
	}
}

func TestGetBookBy_Name_DoesntExists(t *testing.T) {
	books, err := getBook("name", "doesnt_exists")

	if err != nil {
		t.Errorf("Function returned an error, got %s", err)
	}

	if len(books) != 0 {
		t.Errorf("Returned array contains elems")
	}
}

func TestGetBookBy_Isbn_DoesntExists(t *testing.T) {
	books, err := getBook("isbn", "doesnt_exists")

	if err != nil {
		t.Errorf("Function returned an error, got %s", err)
	}

	if len(books) != 0 {
		t.Errorf("Returned array contains elems")
	}
}

func TestBookBy_ErrorCriteria(t *testing.T) {
	books, err := getBook("error", "error")

	if err == nil {
		t.Errorf("Function must return an error, got %s", err)
	}

	if len(books) != 0 {
		t.Errorf("Returned array contains elems")
	}
}
