package main

import (
    "os"
    "testing"
    "net/http"
	  "net/http/httptest"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize()

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request, app App) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}


func TestPingEndpoint(t *testing.T) {
    req, _ := http.NewRequest("GET", "/ping", nil)
    response := executeRequest(req, a)
    checkResponseCode(t, http.StatusOK, response.Code)
}
