package main
import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"net/http"
)

/*
	take url
	generate shortened hash id
	store in url db
	return shortened hash

	when someone visits my url+hash
	serve them the longer url 
*/

// url -> encrypt to shorter id
// 		|-> link-legal characters
//			|-> broad range of characters, short length  

func main()  {
	r := gin.Default()
	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"": "",
		})
	})
	r.Run()
}