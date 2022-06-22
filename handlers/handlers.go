package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func StrangerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "stranger.html", gin.H{})
}
