package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Any("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "Hello,Squirrel",
			"data": "https://github.com/boychai/squirrel-ce",
		})
	})
	r.Run()
}
