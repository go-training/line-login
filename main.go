package main

import (
	"net/http"

	"github.com/go-training/line-login/config"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	conf := config.MustLoad()
	r := setupRouter()
	r.Run(":" + conf.HTTP.Port)
}
