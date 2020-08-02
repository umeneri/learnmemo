package auth

import (
	"fmt"
	"log"

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

func GetUser(c *gin.Context) (goth.User, error) {
	session, _ := store.Get(c.Request, cookieName)
	user, ok := session.Values[userKey].(goth.User)

	if !ok {
		err := fmt.Errorf("cannot get session value")
		return user, err
	}

	return user, nil
}

func SaveSession(user goth.User, c *gin.Context) {
	log.Println(c.Request)
	log.Println(c.Writer)
	session, _ := store.Get(c.Request, cookieName)
	session.Values["authenticated"] = user
	session.Save(c.Request, c.Writer)
}
