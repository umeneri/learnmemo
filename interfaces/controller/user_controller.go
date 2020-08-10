package controller

import (
	"api/interfaces/auth"
	"api/usecase"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/lunny/log"
	"github.com/gin-gonic/gin"
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
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
	  baseURL = "http://localhost:8080"
	}
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), fmt.Sprintf("%s/api/user/callback/google", baseURL)),
	)
}

func (t *userController) Index(c *gin.Context) {
	user, err := auth.GetUser(c)

	if user == nil || err != nil {
		log.Println("Error: user not found")
		log.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}

	if os.Getenv("ENV") == "dev" {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/")
	} else {
		c.HTML(http.StatusOK, "index.html", "")
	}
}

func (t *userController) Entering(c *gin.Context) {
	user, err := auth.GetUser(c)

	if user == nil || err != nil {
		log.Println("Error: user not found")
		log.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
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
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		user := t.userUseCase.FindByEmail(gothUser.Email)
		if user != nil {
			log.Println("Error: user not found")
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		auth.SaveSession(user, c)
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

	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		log.Println("Error: not complete user auth")
		log.Println(err)
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	if user := t.userUseCase.FindByEmail(gothUser.Email); user != nil {
		auth.SaveSession(user, c)
		redirectTo(c, "")
	} else {
		user, err := t.userUseCase.SaveUser(convertToSocialLoginUser(gothUser))
		log.Println("user is ...")
		log.Println(user)
		if user == nil || err != nil {
			log.Println("Error: not complete user save")
			log.Println(err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		err = auth.SaveSession(user, c)
		if err != nil {
			log.Println("Error: not complete user auth")
			log.Println(err)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
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

	user, err := auth.GetUser(c)

	if user == nil || err != nil {
		c.String(http.StatusBadRequest, "Bad request: user not authorized")
		return
	}

	user.Name = userForm.Name

	err = t.userUseCase.UpdateUser(user)
	if err != nil {
		log.Fatal(err)
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
