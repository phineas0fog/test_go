package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type Book struct {
	Name string `json:"name"`
	Isbn string `json:"isbn"`
}

//////////////////////////////////////////////////
func getEM(code int) string {
	if code == http.StatusNotFound {
		return "This book does not exists"
	}
	return "UNKNOWN"
}

func filter(bs []Book, f func(Book) bool) []Book {
	bsf := make([]Book, 0)
	for _, b := range bs {
		if f(b) {
			bsf = append(bsf, b)
		}
	}
	return bsf
}

//////////////////////////////////////////////////

func getBooks() []Book {
	// return a stubbed list of []Book

	return []Book{
		Book{"Lord of the rings", "lotr42"},
		Book{"H2G2", "h2g242"}}

}

func getBook(crit string, value string) ([]Book, error) {
	// get a specified Book

	if crit == "name" {
		return filter(getBooks(), func(b Book) bool {
			return b.Name == value
		}), nil
	} else if crit == "isbn" {
		return filter(getBooks(), func(b Book) bool {
			return b.Isbn == value
		}), nil
	} else {
		return make([]Book, 0), errors.New("Criteria must be \"name\" or \"isbn\"")
	}
}

func setupRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/books", func(c *gin.Context) {
		books := getBooks()

		c.JSON(http.StatusOK, books)
	})

	router.GET("/books/:value", func(c *gin.Context) {
		value := c.Param("value")
		crit := c.Query("crit")

		if len(c.Request.URL.Query()) == 0 {
			crit = "isbn"
		}

		books, err := getBook(crit, value)

		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "If present, criteria must be 'name' or 'isbn'"})
		} else {
			c.JSON(http.StatusOK, books)
		}

	})

	return router
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := setupRouter()
	router.Run(":" + port)
}
