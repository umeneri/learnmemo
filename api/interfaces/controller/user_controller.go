package controller

import (
	"api/interfaces/auth"
	"api/usecase"
	"context"
	"fmt"
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
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:3030/api/user/callback/google"),
	)
}

func (t *userController) Index(c *gin.Context) {
	gothUser, err := auth.GetUser(c)
	user := t.userUseCase.FindByEmail(gothUser.Email)

	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"user": user,
	})
}

func (t *userController) LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}

func (t *userController) Login(c *gin.Context) {
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		auth.SaveSession(user, c)
		t.userUseCase.SaveUser(convertToSocialLoginUser(user))
		c.Redirect(http.StatusTemporaryRedirect, "/")
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
	t.userUseCase.SaveUser(convertToSocialLoginUser(user))
	redirectTo(c, "")
}

func (t *userController) Logout(c *gin.Context) {
	auth.DeleteSession(c)
	redirectTo(c, "login")
}

func redirectTo(c *gin.Context, location string)  {
	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://localhost:3000/%s", location))
	} else {
		c.Redirect(http.StatusTemporaryRedirect, 	fmt.Sprintf("/%s", location))
	}
}

func contextWithProviderName(c *gin.Context, provider string) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}

func convertToSocialLoginUser(user goth.User) usecase.SocialLoginUser {
	return usecase.SocialLoginUser{
		UserID:    user.UserID,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
}
