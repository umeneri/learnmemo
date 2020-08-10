package auth

import (
	"api/domain/model"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	// キーの長さは 16, 24, 32 バイトのいずれかでなければならない。
	// (AES-128, AES-192 or AES-256)
	key        = []byte("super-secret-key")
	store      = sessions.NewCookieStore(key)
	cookieName = "taskball"
	userKey    = "auth"
)

func init()  {
	gob.Register(model.User{})
}

func AuthRequired(c *gin.Context) {
	log.Println("with session")

	user, err := GetUser(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	} else {
		c.Next()
	}
}

func AuthRequiredPage(c *gin.Context) {
	log.Println("with session")

	user, err := GetUser(c)
	if user == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else {
		c.Next()
	}
}

func GetUser(c *gin.Context) (*model.User, error) {
	session, _ := store.Get(c.Request, cookieName)
	user, ok := session.Values[userKey].(model.User)

	log.Println(user)

	if !ok {
		err := fmt.Errorf("cannot get session value")
		return &user, err
	}

	return &user, nil
}

func GetUserId(c *gin.Context) (int64, error) {
	session, _ := store.Get(c.Request, cookieName)
	user, ok := session.Values[userKey].(model.User)

	log.Println(user)

	if !ok {
		err := fmt.Errorf("cannot get session value")
		return 0, err
	}

	return user.Id, nil
}

func SaveSession(user *model.User, c *gin.Context) error {
	session, _ := store.Get(c.Request, cookieName)
	session.Values[userKey] = user
	err := session.Save(c.Request, c.Writer)
	return err
}

func DeleteSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, cookieName)
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)
	return err
}
