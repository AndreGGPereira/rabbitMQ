package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Decode os dados do body
func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

//Pega o parametro uuid enviados na request
func getParameter(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["uuid"]
}

//Envia a resposta da request
func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}

//timeNowFormatted pegar a data atual e formata ISO 8601
func timeNowFormatted() string {
	const layoutDate = "2006-01-02T15:04:05.999999"
	return time.Now().Format(layoutDate)
}
