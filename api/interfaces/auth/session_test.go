package auth_test

import (
	"api/interfaces/auth"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

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
	auth.SaveSession(user, c)

	resopnseHeader := c.Writer.Header()
	str := resopnseHeader.Get("Set-Cookie")

	if str == "" {
		t.Fatalf("cannot get session value")
	}
}

func TestDeleteSession(t *testing.T) {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	auth.DeleteSession(c)

	resopnseHeader := c.Writer.Header()
	str := resopnseHeader.Get("Set-Cookie")
	fmt.Println(str)

	if strings.Index(str, "Max-Age=0") == -1 {
		t.Fatalf("session value is not deleted")
	}
}
