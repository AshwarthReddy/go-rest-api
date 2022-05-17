package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// const db := make([]string)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: 1, Name: "UPDATE YOUR JAVA", Author: "Aswarth"},
	{ID: 2, Name: "LEARN-GOLANG", Author: "some"},
	{ID: 3, Name: "Hadling stress", Author: "xyb"},
}

func main() {

	router := gin.New()

	// defult GET METHOD
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	// get all books
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	// fetch book based on id
	router.GET("/:id", func(c *gin.Context) {

		idString := c.Param("id")
		id, _ := strconv.Atoi(idString)

		for index, book := range books {
			fmt.Println(index, ":", book)
			if book.ID == id {
				c.JSON(http.StatusOK, &book)
				break
			}
		}
	})

	// save book
	router.POST("/", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
		c.JSON(http.StatusAccepted, book)

	})

	// update book

	router.PUT("/update-book", func(c *gin.Context) {

		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {

			fmt.Printf("book id %d book name %s, book author %s", book.ID, book.Name, book.Author)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Printf("book id %d book name %s, book author %s", book.ID, book.Name, book.Author)

		for _, data := range books {
			if data.ID == book.ID {
				fmt.Println("inside if condition")
				data = Book{data.ID, book.Name, book.Author}
				books = append(books, data)
				c.JSON(http.StatusAccepted, book)
				break
			}
		}

	})

	// delete book
	router.DELETE("/:id", func(c *gin.Context) {

		strId := c.Param("id")

		if id, err := strconv.Atoi(strId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {

			for index, book := range books {

				if book.ID == id {
					books = append(books[:index], books[index+1:]...)
					break
				}
			}
		}
		c.JSON(http.StatusAccepted, books)

	})

	router.Run("localhost:9090")

}
