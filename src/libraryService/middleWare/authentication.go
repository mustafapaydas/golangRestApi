package middleWare

import (
	"LibraryApi/src/libraryService/business"
	"context"
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

type SessionScope struct {
	authorizations []string
	firstName      string
	lastName       string
	token          string
}

var SessionAuthCodeList = []string{}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if strings.Contains(tokenString, "Bearer") {
			tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		if !checkToken(tokenString) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UnAuthorized",
			})
			//throw 403
			return
		}
		c.Next()
	}
}

func AuthCheck(_auth string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("Success Auth", _auth)
		if !HasAuth(_auth) {
			ctx.AbortWithStatus(403)
			//TODO:return specific error html
			ctx.HTML(http.StatusForbidden, "errors/error.tmpl", gin.H{
				"message": "403 errors",
			})
			return
		}

		ctx.Next()
	}
}

func HasAuth(_auth string) bool {
	for _, _code := range SessionAuthCodeList {
		if _auth == _code {
			return true
		}
	}
	return false
}

func checkToken(userToken string) bool {
	jwksURL := "http://localhost:5001/realms/library/protocol/openid-connect/certs"
	ctx, cancel := context.WithCancel(context.Background())
	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}
	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}
	jwtB64 := userToken
	token, err := jwt.Parse(jwtB64, jwks.Keyfunc)
	if err != nil {
		log.Println("Failed to parse the JWT.\nError: %s", err.Error())
	}
	if !token.Valid {
		log.Println("The token is not valid.")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		//log.Println(fmt.Sprintf("%v succeded", claims["sid"]))
		//log.Println(fmt.Sprintf("%v succeded", claims["groups"]))
		groupClaims := claims["groups"].([]interface{})
		roleList := []string{}
		for i := range claims["groups"].([]interface{}) {
			roleList = append(roleList, groupClaims[i].(string))
		}

		if claims["groups"] != nil {
			fmt.Println(business.FindByName(roleList))
			for _, code := range business.FindByName(roleList).Authorizations {
				SessionAuthCodeList = append(SessionAuthCodeList, code.Code)
			}
			//SessionAuthCodeList
			fmt.Println(SessionAuthCodeList)
		}

	}

	cancel()
	jwks.EndBackground()
	if token.Valid {
		return true
	} else {
		return false
	}
}

//
//func CCCC() {
//	book := new(model.Book)
//	book.Name = "cccccc"
//	cat := new(model.Category)
//	cat.Id = 1
//	book.Category = *cat
//	_s := business.AbstractService{}
//	_s.Create(&book)
//	fmt.Println(book)
//}
