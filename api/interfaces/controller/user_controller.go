package controller

import (
	"api/usecase"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/lunny/log"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type UserController interface {
	Index(c *gin.Context)
	LoginIndex(c *gin.Context)
	Login(c *gin.Context)
	Callback(c *gin.Context)
}

type userController struct {
	userUseCase usecase.UserUseCase
}

var (
	// キーの長さは 16, 24, 32 バイトのいずれかでなければならない。
	// (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
	cookieName = "taskball"
)

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
	session, _ := store.Get(c.Request, cookieName)

	if user, ok := session.Values["authenticated"].(goth.User); !ok || user.UserID == "" {
		log.Println(user)
		c.String(http.StatusForbidden, "Forbidden")
	} else {
		log.Println(user)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user": user,
		})
	}
}

func (t *userController) LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}

func (t *userController) Login(c *gin.Context) {
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		log.Println(gothUser)
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

	log.Println(user)

	session, _ := store.Get(c.Request, cookieName)
	session.Values["authenticated"] = user
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func contextWithProviderName(c *gin.Context, provider string) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
}
