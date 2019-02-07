package main

import (
	"net/http"

	"github.com/go-training/line-login/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func setupRouter() *gin.Engine {
	conf := config.MustLoad()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Static("/images", "./images")
	r.LoadHTMLGlob("templates/*")

	// user login page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":    "Line QR Code Login Example",
			"lineID":   conf.Line.ID,
			"callback": conf.Line.Callback,
		})
	})

	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		changed := c.Query("friendship_status_changed")
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"state":   state,
			"changed": changed,
		})
	})

	return r
}

func main() {
	conf := config.MustLoad()
	r := setupRouter()
	r.Run(":" + conf.HTTP.Port)
}
