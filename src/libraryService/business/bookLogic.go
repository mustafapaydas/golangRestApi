package business

import (
	"LibraryApi/src/libraryService/db"
	"LibraryApi/src/libraryService/model"
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
)

var conn = db.ConnectToDb()
var service AbstractService

type BookLogic struct {
	*AbstractService
}

func bind(book model.Book, newBook model.Book) model.Book {
	newBook.Name = book.Name
	newBook.Isbn = book.Isbn
	newBook.PageCount = book.PageCount
	newBook.Count = book.Count
	return newBook
}

// @Router /bookapi/book [put]
func Update(c *gin.Context) {
	book := model.Book{}
	c.BindJSON(&book)
	conn.Model(model.Book{}).Where("id = ?", book.Id).Updates(&book)
	c.JSON(http.StatusOK, &book)
}

// @Param        book  body      model.Book  true  "Book JSON"
// @Router /bookapi/book [post]
func CreateBok(c *gin.Context) {
	book := model.Book{}
	err := c.BindJSON(&book)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if book.Id != 0 {
		var NotFoundError = fmt.Errorf("Id Must Be Null")
		c.JSON(http.StatusInternalServerError, NotFoundError)
		return
	}
	var newBook model.Book
	newBook = bind(book, newBook)
	if result := conn.Create(&newBook); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &book)
}

// @Router /bookapi/books [get]
func FindAll(c *gin.Context) {
	var books []model.Book
	c.BindJSON(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// @Param  id   /bookapi/find     int  true  "get book by id"
// @Router /bookapi/find [get]
func Find(c *gin.Context) {
	var book model.Book

	//if err := conn.Find(&book).Error; err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//	return
	//}
	conn.First(&book, c.Param("id"))
	c.JSON(200, book)
}
