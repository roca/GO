package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"udemy.com/Frameworks/Gin-Web/books"
)

func main() {
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// the hello message endpoint with JSON response from map
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework."})
	})

	// get all books
	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books.AllBooks())
	})

	// create new book
	engine.POST("/api/books", func(c *gin.Context) {
		var book books.Book
		if c.BindJSON(&book) == nil {
			isbn, created := books.CreateBook(book)
			if created {
				c.Header("Location", "/api/books/"+isbn)
				c.Status(http.StatusCreated)
			} else {
				c.Status(http.StatusConflict)
			}
		}
	})

	// get book by ISBN
	engine.GET("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		book, found := books.GetBook(isbn)
		if found {
			c.JSON(http.StatusOK, book)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	// update existing book
	engine.PUT("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")

		var book books.Book
		if c.BindJSON(&book) == nil {
			exists := books.UpdateBook(isbn, book)
			if exists {
				c.Status(http.StatusOK)
			} else {
				c.Status(http.StatusNotFound)
			}
		}
	})

	// delete book
	engine.DELETE("/api/books/:isbn", func(c *gin.Context) {
		isbn := c.Params.ByName("isbn")
		books.DeleteBook(isbn)
		c.Status(http.StatusOK)
	})

	// configuration for static files and templates
	engine.LoadHTMLGlob("./templates/*.html")
	engine.StaticFile("/favicon.ico", "./favicon.ico")

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Advanced Cloud Native Go with Gin Framework",
		})
	})

	// run server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
