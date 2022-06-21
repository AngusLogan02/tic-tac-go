package main

import (
	"tic-tac-go/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("public/html/*")
	r.Static("/static/css", "./public/css")
	r.Static("/static/js", "./public/js")

	r.GET("/", handlers.IndexHandler)

	r.Run()
}
