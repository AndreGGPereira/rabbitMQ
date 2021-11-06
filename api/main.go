package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	c := cors.AllowAll()
	router.HandleFunc("/cliente", CreateCliente).Methods("POST")
	router.HandleFunc("/cliente/{uuid}", GetCliente).Methods("GET")
	router.HandleFunc("/clientes", GetClientes).Methods("GET")
	router.HandleFunc("/cliente", UpdateCliente).Methods("PUT")
	router.HandleFunc("/cliente/{uuid}", DeleteCliente).Methods("DELETE")
	handler := c.Handler(router)
	if os.Getenv("PORT") == "" {
		http.ListenAndServe(":8080", handler)
	} else {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Println("Port was defined but could not be parsed.")
			os.Exit(1)
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
	}

}
