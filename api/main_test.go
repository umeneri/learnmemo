// file: main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func teardown(ts *httptest.Server) {
    ts.Close()
}

func TestTaskList(t *testing.T) {
    ts := httptest.NewServer(setupServer())
    defer teardown(ts)

    resp, err := http.Get(fmt.Sprintf("%s/task/v1/list", ts.URL))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    checkResponseHeader(t, resp, 200)
}

func TestTaskAdd(t *testing.T) {
    ts := httptest.NewServer(setupServer())
    defer teardown(ts)

    requestBody, err := json.Marshal(map[string]interface{}{
        "title": "title1",
        "progressMinute": 10,
         "status": 1,
    })

    resp, err := http.Post(fmt.Sprintf("%s/task/v1/add", ts.URL), "application/json", bytes.NewBuffer(requestBody))

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    checkResponseHeader(t, resp, 201)
}

func checkResponseHeader(t *testing.T, resp *http.Response, statusCode int)  {
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
