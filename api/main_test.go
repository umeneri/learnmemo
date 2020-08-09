// file: main_test.go

package main

import (
	"api/domain/model"
	"api/infrastructure/repository"
	"api/interfaces/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
)

var ts *httptest.Server
var userRepository = repository.NewUserRepository("gin_test")
var testUser = model.User{
		Email: "hoge1@gmail.com",
		Name:       "hoge1",
		ProviderId: "hoge1",
		AvatarUrl:  "hoge1",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func TestListTask(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/api/task/v1/list", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 200)
}

func TestAddTask(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"title":       "title1",
		"elapsedTime": 10,
		"status":      1,
	})

	resp, err := http.Post(fmt.Sprintf("%s/api/task/v1/add", ts.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 201)
}

func TestUserIndex(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedCode := 200
	if resp.StatusCode != expectedCode {
		t.Fatalf("Expected status code %d, got %v", expectedCode, resp.StatusCode)
	}
}

func saveSessionAndUser() *http.Cookie {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	gothUser := goth.User{
		Email: "hoge1@gmail.com",
	}
	auth.SaveSession(gothUser, c)

	_, err := userRepository.SaveUser(&testUser)
	if err != nil {
		return nil
	}

	parser := &http.Request{Header: http.Header{"Cookie": c.Writer.Header()["Set-Cookie"]}}
	taskball, _ := parser.Cookie("taskball")
	return taskball
}

func TestUpdateUser(t *testing.T) {
	taskball := saveSessionAndUser()
	if taskball == nil {
		t.Fatalf("user was not saved")
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"name": "name1",
	})

	url := fmt.Sprintf("%s/api/user/v1/update", ts.URL)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	req.AddCookie(taskball)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedCode := 200
	if resp.StatusCode != expectedCode {
		t.Fatalf("Expected status code %d, got %v", expectedCode, resp.StatusCode)
	}
}

func setup() {
	engine := setupServer("test")
	ts = httptest.NewServer(engine)
}

func teardown() {
	err := userRepository.DeleteUser(&testUser)
	fmt.Println(err)
	ts.Close()
}

func checkResponseHeader(t *testing.T, resp *http.Response, statusCode int) {
	if resp.StatusCode != statusCode {
		t.Fatalf("Expected status code %d, got %v", statusCode, resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	// Assert that the "content-type" header is actually set
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	// Assert that it was set as expected
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
