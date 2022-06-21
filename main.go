package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("public/html")
	r.Static("/static/css", "public/css")
}
