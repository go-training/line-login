package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-training/line-login/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/resty.v1"
)

// AuthSuccess get access token
type AuthSuccess struct {
	AccessToken  string  `json:"access_token"`
	ExpiresIn    float64 `json:"expires_in"`
	IDToken      string  `json:"id_token"`
	RefreshToken string  `json:"refresh_token"`
	Scope        string  `json:"scope"`
	TokenType    string  `json:"token_type"`
}

// Profile Gets a user's display name, profile image, and status message.
// See: https://developers.line.biz/en/reference/social-api/#get-user-profile
type Profile struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
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
		resp, _ := resty.R().
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
		fmt.Printf("\nResponse Body: %v", resp)

		profile := &Profile{}
		resp, _ = resty.R().
			SetHeader("Authorization", "Bearer "+authSuccess.AccessToken).
			SetResult(profile). // or SetResult(AuthSuccess{}).
			Get("https://api.line.me/v2/profile")
		fmt.Printf("\nResponse Body: %v", resp)

		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":         "Line QR Code Login Example",
			"userID":        profile.UserID,
			"displayName":   profile.DisplayName,
			"pictureURL":    profile.PictureURL,
			"statusMessage": profile.StatusMessage,
		})
		return
	})

	return r
}

func main() {
	log.Fatal(RunHTTPServer(context.Background()))
}
