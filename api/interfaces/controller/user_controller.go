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
	"github.com/k0kubun/pp"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type UserController interface {
	Index(c *gin.Context)
	Entering(c *gin.Context)
	LoginIndex(c *gin.Context)
	Login(c *gin.Context)
	Callback(c *gin.Context)
	Logout(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

type UserForm struct {
	Name string
}

func NewUserController(useCase usecase.UserUseCase) UserController {
	return &userController{
		userUseCase: useCase,
	}
}

func init() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/api/user/callback/google"),
	)
}

func (t *userController) Index(c *gin.Context) {
	gothUser, err := auth.GetUser(c)
	user := t.userUseCase.FindByEmail(gothUser.Email)

	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

func (t *userController) Entering(c *gin.Context) {
	gothUser, err := auth.GetUser(c)
	user := t.userUseCase.FindByEmail(gothUser.Email)

	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/entering")
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

func (t *userController) LoginIndex(c *gin.Context) {
	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/login")
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
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

	if t.userUseCase.FindByEmail(user.Email) != nil {
		redirectTo(c, "")
	} else {
		t.userUseCase.SaveUser(convertToSocialLoginUser(user))
		redirectTo(c, "entering")
	}
}

func (t *userController) Logout(c *gin.Context) {
	auth.DeleteSession(c)
	redirectTo(c, "login")
}

func (t *userController) UpdateUser(c *gin.Context) {
	userForm := UserForm{}
	err := c.BindJSON(&userForm)

	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, "Bad request: invalid form")
		return
	}

	gothUser, err := auth.GetUser(c)

	pp.Println(gothUser)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request: user not authorized")
		return
	}

	user := t.userUseCase.FindByEmail(gothUser.Email)
	pp.Println(user)
	if user == nil {
		c.String(http.StatusBadRequest, "Bad request: user not authorized")
		return
	}

	user.Name = userForm.Name

	err = t.userUseCase.UpdateUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func redirectTo(c *gin.Context, location string) {
	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("http://localhost:3000/%s", location))
	} else {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/%s", location))
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
