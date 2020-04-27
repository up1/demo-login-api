package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccessWithUserNameAndPassword(t *testing.T) {
	var request = []byte(`{"username":"Somkiat", "password": "PasswordSomkiat"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(request))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"firstname":"Somkiat","lastname":"Puisung","email":"somkiat@xxx.com"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
