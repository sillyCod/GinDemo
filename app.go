package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Love struct {
	Name string `form:"name" binding:"required"`
	Time string `form:"time" binding:"required"`
}

func main() {
	app := gin.Default()
	app.GET("/love", func(context *gin.Context) {
		love:=new(Love)
		context.BindQuery(love)
		fmt.Println(love.Name, love.Time)
		context.JSON(http.StatusOK, gin.H{"love": "ww"})
	})
	app.Use(myMiddleware)
	app.Handle("GET", "/definitely", func(context *gin.Context) {
		//context.HTML(http.StatusOK, "I love ww", "I love ww")
		status, err := context.Writer.Write([]byte("I love ww"))
		if status == http.StatusOK {
			fmt.Println("success")
		}
		if err != nil {
			fmt.Println(err)
		}
		//context.PureJSON()
	})

	app.Run(":8888")
}

func myMiddleware(ctx *gin.Context) {
	fmt.Println("My middleware called")
	ctx.Next()
}
