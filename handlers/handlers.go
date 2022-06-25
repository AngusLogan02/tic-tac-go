package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"roomLink": RandomString(8)})
}

func FriendHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "friend.html", gin.H{"friendID": c.Param("friendID")})
}

func StrangerHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "stranger.html", gin.H{})
}
