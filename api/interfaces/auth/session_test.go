package auth_test

import (
	"api/interfaces/auth"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
)

func TestSaveSession(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	user := goth.User{
		Email: "hoge@gmail.com",
	}
	fmt.Println(c)
	fmt.Println(user)
	auth.SaveSession(user, c)

	resopnseHeader := c.Writer.Header()
	str := resopnseHeader.Get("Set-Cookie")
	fmt.Println(str)

	if str == "" {
		t.Fatalf("cannot get session value")
	}
}
