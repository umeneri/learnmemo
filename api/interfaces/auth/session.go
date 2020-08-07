package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
)

var (
	// キーの長さは 16, 24, 32 バイトのいずれかでなければならない。
	// (AES-128, AES-192 or AES-256)
	key        = []byte("super-secret-key")
	store      = sessions.NewCookieStore(key)
	cookieName = "taskball"
	userKey    = "auth"
)

func GetUserId() int64 {
	return 1
}

func AuthRequired(c *gin.Context) {
	log.Println("with session")

	user, err := GetUser(c)
	if user.UserID == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	} else {
		c.Next()
	}
}

func GetUser(c *gin.Context) (goth.User, error) {
	session, _ := store.Get(c.Request, cookieName)
	user, ok := session.Values[userKey].(goth.User)

	log.Println(user)

	if !ok {
		err := fmt.Errorf("cannot get session value")
		return user, err
	}

	return user, nil
}

func SaveSession(user goth.User, c *gin.Context) {
	session, _ := store.Get(c.Request, cookieName)
	session.Values[userKey] = user
	session.Save(c.Request, c.Writer)
}

func DeleteSession(c *gin.Context) {
	session, _ := store.Get(c.Request, cookieName)
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
}
