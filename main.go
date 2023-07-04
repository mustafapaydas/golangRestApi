package main

import (
	"LibraryApi/src/libraryService/api"
	"LibraryApi/src/libraryService/db"
	"LibraryApi/src/libraryService/middleWare"
	"github.com/gin-gonic/gin"
)

func main() {
	//middleWare.CCCC()

	db.ConnectToDb()

	router := gin.Default()
	v1 := router.Group("/api/v1", middleWare.AuthMiddleware())
	api.SetupBookRoutes(v1)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8070")
}
