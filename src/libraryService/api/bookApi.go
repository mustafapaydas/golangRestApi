package api

import (
	"LibraryApi/src/libraryService/business"
	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(router *gin.RouterGroup) {

	book := router.Group("/book")
	book.GET("/books", business.FindAll)
	book.POST("/book", business.CreateBok)
	book.PUT("/update", business.Update)
}

func getUserList(c *gin.Context) {
	// Kullanıcı listesi alımı işlemleri
}

func createUser(c *gin.Context) {
	// Kullanıcı oluşturma işlemleri
}

func getUser(c *gin.Context) {
	// Belirli bir kullanıcıyı alma işlemleri
}

func updateUser(c *gin.Context) {
	// Kullanıcı güncelleme işlemleri
}

func deleteUser(c *gin.Context) {
	// Kullanıcı silme işlemleri
}
