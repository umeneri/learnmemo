package main

import (
	"api/domain/model"
	"api/infrastructure/repository"
	"api/interfaces/auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var ts *httptest.Server
var userRepository = repository.NewUserRepository("gin_test")
var testUser = model.User{
	Email:      "hoge1@gmail.com",
	Name:       "hoge1",
	ProviderId: "hoge1",
	AvatarUrl:  "hoge1",
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
}
var cookie *http.Cookie

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func TestAddTask(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"title":       "title1",
		"elapsedTime": 10,
		"status":      1,
	})
	url := fmt.Sprintf("%s/api/task/v1/add", ts.URL)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	req.AddCookie(cookie)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 201)
}

type ListTaskResponse struct {
	Data    []model.Task
	Message string
}

func TestListTask(t *testing.T) {
	resp, err := httpGet(t, "api/task/v1/list")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 200)

	defer resp.Body.Close()
	var taskResp ListTaskResponse
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &taskResp)

	if len(taskResp.Data) == 0 {
		t.Fatalf("error: tasks is empty")
	}
}

func TestListTaskNotAuthenticated(t *testing.T) {
	url := fmt.Sprintf("%s/api/task/v1/list", ts.URL)
	resp, err := http.Get(url)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 401)
}

func TestUserIndex(t *testing.T) {
	resp, err := httpGet(t, "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedCode := 200
	if resp.StatusCode != expectedCode {
		t.Fatalf("Expected status code %d, got %v", expectedCode, resp.StatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"name": "name1",
	})
	url := fmt.Sprintf("%s/api/user/v1/update", ts.URL)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	req.AddCookie(cookie)
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
	saveSessionAndUser()
}

func saveSessionAndUser() {
	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("GET", "/", nil)
	auth.SaveSession(&testUser, c)

	_, err := userRepository.SaveUser(&testUser)
	if err != nil {
		fmt.Println("Error in save user")
		os.Exit(1)
	}

	parser := &http.Request{Header: http.Header{"Cookie": c.Writer.Header()["Set-Cookie"]}}
	cookie, err = parser.Cookie("taskball")

	if err != nil {
		fmt.Println("Error in parse cookie")
		os.Exit(1)
	}
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

func httpGet(t *testing.T, path string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", ts.URL, path)
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
	if err != nil {
		t.Fatalf("create get repuest error")
	}

	req.AddCookie(cookie)
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
