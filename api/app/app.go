package app

import (
	"api/app/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartService() {

	app := New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)
	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)
	runServer()
}

func runServer() {
	log.Println("App running..")
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

type App struct {
	Router *mux.Router
	DB     database.InterfaceClientDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
}

//Instanciação das rotas
func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/client", a.CreateClientHandler()).Methods("POST")
	a.Router.HandleFunc("/client/{uuid}", a.UpdateClientHandler()).Methods("PUT")
	a.Router.HandleFunc("/client/{uuid}", a.GetClientByIDHandler()).Methods("GET")
	a.Router.HandleFunc("/clients", a.GetClientsHandler()).Methods("GET")
	a.Router.HandleFunc("/client/{uuid}", a.DeleteClientHandler()).Methods("DELETE")

}

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
