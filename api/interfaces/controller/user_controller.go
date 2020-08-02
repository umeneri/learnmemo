package controller

import (
	"api/usecase"
	"net/http"
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gin-gonic/gin"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

type UserController interface {
	Index(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: useCase,
	}
}

func (t *userController) Index(c *gin.Context) {
	// c.JSON(http.StatusCreated, gin.H{
	// 	"status": "ok",
	// })
	c.HTML(http.StatusOK, "index.html", "")
}

func (t *userController) Login(c *gin.Context) {
}