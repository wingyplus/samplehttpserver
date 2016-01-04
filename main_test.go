package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex_GET(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://1.2.3.4/hello", nil)

	index(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expect status 200 OK but got", recorder.Code)
	}
	if body := recorder.Body.String(); body != "Hello World with GET" {
		t.Error("Expect response body 'Hello World' but got", body)
	}
}

func TestIndex_POST(t *testing.T) {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "http://1.2.3.4/hello", strings.NewReader(`{"name":"blah"}`))

	index(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Error("Expect status 200 OK but got", recorder.Code)
	}
	if body := recorder.Body.String(); body != "Hello World - blah" {
		t.Error("Expect response body 'Hello World - blah' but got", body)
	}
}

func TestIndex_MethodNotAllowed(t *testing.T) {
	for _, method := range []string{"PUT", "DELETE", "HEAD", "OPTION"} {
		recorder := httptest.NewRecorder()
		req, _ := http.NewRequest(method, "http://1.2.3.4/hello", nil)

		index(recorder, req)

		if recorder.Code != http.StatusMethodNotAllowed {
			t.Error("Expect status 405 Method Not Allowed but got", recorder.Code)
		}
	}
}
