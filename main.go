package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"libraryApi/src/libraryService/business"
	"libraryApi/src/libraryService/db"
	"libraryApi/src/libraryService/model"
)

func main() {
	db.ConnectToDb()

	//gin'in varsayılan ayarlarında bir yönlendirici oluşturalım.
	router := gin.Default()

	//anasayfayı inde fonksiyonumuz yakalayacak
	router.GET("/", index)
	router.GET("/books", business.FindAll)
	router.POST("/book", business.CreateBok)
	router.PUT("/update", business.Update)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/find/:id", business.Find)

	//daha sonra sunucuyu başlatıyoruz
	router.Run(":8070")
}

func generate(c *gin.Context) {
	var book model.Book
	var d = c.ShouldBind(book)
	if c.ShouldBind(&book) == nil {
		fmt.Println(book.Name)
	}
	fmt.Println(d)
	business.CreateBok(c)

}

// anasayfayı yakalayacak olan fonksiyonumuz
func index(c *gin.Context) {
	//c ile gin nesnemize bağlam oluşturduk.
	//c'yi kullanarak artık gin özelliklerine erişebiliriz.

	//sayfaya düz yazı gönderdik
	c.String(200, "Merhaba Dünya")
	//Buradaki 200 sunucudan bir cevap geldiğini anlamına gelir
}
