package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Implementar testes
func TestCreateClientHandler(t *testing.T) {

}
func TestUpdateClientHandler(t *testing.T) {
}

func TestGetClientsHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/clients", nil)
	w := httptest.NewRecorder()
	upperCaseHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) == "" {
		t.Errorf("expected nil got %v", string(data))

	}
}

func upperCaseHandler(w http.ResponseWriter, r *http.Request) {
}

func TestDeleteClientHandler(t *testing.T) {}
