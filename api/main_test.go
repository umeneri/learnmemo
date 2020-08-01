// file: main_test.go

package main

import (
	"api/interfaces/controller"
	"api/interfaces/router"
	"api/service"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
    setup()
    ret := m.Run()
    teardown()
    os.Exit(ret)
}

func TestTaskList(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/task/v1/list", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 200)
}

func TestTaskAdd(t *testing.T) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"title":          "title1",
		"progressMinute": 10,
		"status":         1,
	})

	resp, err := http.Post(fmt.Sprintf("%s/task/v1/add", ts.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	checkResponseHeader(t, resp, 201)
}

func setup() {
	if ts == nil {
		dbName := "gin_test"
		taskService := service.NewTaskService(dbName)
		taskController := controller.NewTaskController(taskService)
		engine := router.SetupRoute(taskController)
		ts = httptest.NewServer(engine)
	}
}

func teardown() {
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
