package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//Implementar testes
func TestCreateClientHandler(t *testing.T) {
}

func TestUpdateClientHandler(t *testing.T) {}

func TestGetClientByIDHandler(t *testing.T) {}

func TestGetClientsHandler(t *testing.T) {
}

func TestDeleteClientHandler(t *testing.T) {}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	//a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
