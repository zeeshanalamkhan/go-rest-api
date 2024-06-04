package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zeeshanalamkhan/go-crud-api/controller"
	"github.com/zeeshanalamkhan/go-crud-api/database"
)

func main() {
	fmt.Println("starting application .... ")
	database.DatabaseConncetion()

	r := gin.Default()
	r.GET("/books/:id", controller.ReadBook)
	r.GET("/books", controller.ReadBooks)
	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)
	r.Run(":5000")
}
