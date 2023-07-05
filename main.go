package main

import (
	"LibraryApi/docs"
	"LibraryApi/src/libraryService/api"
	"LibraryApi/src/libraryService/db"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	cmder := exec.Command("swag", "fmt")
	cmder.Dir = basepath
	cmd := exec.Command("swag", "init")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(stdout)
	}
}

func main() {

	docs.SwaggerInfo.BasePath = "/api/v1"
	db.ConnectToDb()
	router := gin.Default()
	url := ginSwagger.URL("http://localhost:8070/swagger/doc.json")

	//router.GET("/xxx", gen)
	v1 := router.Group("/api/v1")
	api.SetupBookRoutes(v1)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":8070")
}
