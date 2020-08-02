package controller

import (
	"api/interfaces/auth"
	"api/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type UserController interface {
	Index(c *gin.Context)
	LoginIndex(c *gin.Context)
	Login(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: useCase,
	}
}

func init() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:3000/user/v1/callback/google"),
	)
}

func (t *userController) Index(c *gin.Context) {
	user, err := auth.GetUser(c)

	log.Println("index")
	log.Println(user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"user": user,
	})
}

func (t *userController) LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}

func (t *userController) Login(c *gin.Context) {
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		auth.SaveSession(user, c)
		c.Redirect(http.StatusTemporaryRedirect, "/user/v1/me")
	} else {
		provider := c.Param("provider")
		c.Request = contextWithProviderName(c, provider)
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (t *userController) Callback(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = contextWithProviderName(c, provider)

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	auth.SaveSession(user, c)

	c.Redirect(http.StatusTemporaryRedirect, "/user/v1/me")
}

func (t *userController) Logout(c *gin.Context) {
	auth.DeleteSession(c)
	c.Redirect(http.StatusTemporaryRedirect, "/user/v1/login")
}

func contextWithProviderName(c *gin.Context, provider string) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}
