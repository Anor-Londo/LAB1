package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHello(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandle)
	handler.ServeHTTP(rr, req)
	resp := rr.Result()
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedCtype := "text/plain"
	if ctype := resp.Header.Get("Content-Type"); ctype != expectedCtype {
		t.Errorf("handler returned unexpected Content-Type: got %v want %v",
			ctype, expectedCtype)
	}

	expectedBody := []byte("Hello, world!")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Compare(body, expectedBody) != 0 {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}

func TestSayHi(t *testing.T) {
	tests := map[string][]byte{
		"/alpha": []byte("Hello, alpha!"),
		"/beta":  []byte("Hello, beta!"),
		"/delta": []byte("Hello, delta!"),
	}

	for url, expectedBody := range tests {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/{name}", sayHi)
		router.ServeHTTP(rr, req)

		resp := rr.Result()
		if status := resp.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned unexpected status code: got %v want %v",
				status, http.StatusOK)
		}

		expectedCtype := "text/plain"
		if ctype := resp.Header.Get("Content-Type"); ctype != expectedCtype {
			t.Errorf("handler returned unexpected Content-Type: got %v want %v",
				ctype, expectedCtype)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(body, expectedBody) != 0 {
			t.Errorf("handler returned unexpected body: got %v want %v", body, expectedBody)
		}
	}
}
