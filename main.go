package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity  int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Title 1", Author: "Author 1", Quantity: 2},
	{ID: "2", Title: "Title 2", Author: "Author 2", Quantity: 10},
	{ID: "3", Title: "Title 3", Author: "Author 3", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}