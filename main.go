package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLFiles("index.tmpl")

	engine.Static("/pages", "./pages")
	engine.Static("/public", "./public")
	engine.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello")
	})

	engine.Use(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	if err := engine.Run(":8010"); err != nil {
		log.Fatalln(err)
		return
	}
}
