package business

import (
	"github.com/gin-gonic/gin"
	"libraryApi/src/libraryService/db"
	"libraryApi/src/libraryService/model"
	"net/http"
)

var conn = db.ConnectToDb()

func CreateBok(c *gin.Context) {
	book := model.Book{}
	if err := c.BindJSON(&book); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var newBook model.Book
	newBook.Name = book.Name
	newBook.Isbn = book.Isbn
	newBook.PageCount = book.PageCount
	if result := conn.Create(&newBook); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &book)
}
func FindAll(c *gin.Context) {
	var books []model.Book
	conn.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
