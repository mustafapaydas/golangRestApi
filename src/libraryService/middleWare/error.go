package middleWare

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//func LogError(c *gin.Context, code int, message string) gin.HandlerFunc {
//
//}
