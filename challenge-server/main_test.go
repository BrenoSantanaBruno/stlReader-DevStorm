// main_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"stlReader-DevStorm/web/handlers"
	"testing"
)

func TestUploadFileHandler(t *testing.T) {
	// Create an HTTP request to test the handler function
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (implements http.ResponseWriter) to capture the response
	rr := httptest.NewRecorder()

	// Create an HTTP handler using the UploadFileHandler function from the handlers package
	handler := http.HandlerFunc(handlers.UploadFileHandler)

	// Call the HTTP handler
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body if necessary
	// if body := rr.Body.String(); body != "Expected" {
	// 	t.Errorf("handler returned wrong body: got %v, want %v", body, "Expected")
	// }
}

func TestMain(m *testing.M) {
	// Run your main function as a test
	go main()

	// You can add additional tests here to check the behavior of the HTTP server
	// using libraries like "net/http/httptest" to create requests and check responses.
}
