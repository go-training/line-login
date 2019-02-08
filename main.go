package main

import (
	"fmt"
	"net/http"

	"github.com/go-training/line-login/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/resty.v1"
)

type AuthSuccess struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	IDToken      string  `json:"id_token"`
	RefreshToken string  `json:"refresh_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}

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
		// state := c.Query("state")
		// changed := c.Query("friendship_status_changed")

		authSuccess := &AuthSuccess{}
		resp, err := resty.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetFormData(map[string]string{
				"grant_type":    "authorization_code",
				"code":          code,
				"redirect_uri":  conf.Line.Callback,
				"client_id":     conf.Line.ID,
				"client_secret": conf.Line.Secret,
			}).
			SetResult(authSuccess). // or SetResult(AuthSuccess{}).
			Post("https://api.line.me/oauth2/v2.1/token")
		fmt.Printf("\nResponse Error: %v", err)
		fmt.Printf("\nResponse Body: %v", resp)

		c.JSON(http.StatusOK, gin.H{
			"AccessToken":  authSuccess.AccessToken,
			"ExpiresIn":    authSuccess.ExpiresIn,
			"IDToken":      authSuccess.IDToken,
			"RefreshToken": authSuccess.RefreshToken,
			"Scope":        authSuccess.Scope,
		})
	})

	return r
}

func main() {
	conf := config.MustLoad()
	r := setupRouter()
	r.Run(":" + conf.HTTP.Port)
}
