package app

import (
	"api/app/models"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}
}

//CreateClientHandler
func (a *App) CreateClientHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Cliente{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		c := &models.Cliente{
			Nome:     req.Nome,
			Endereco: req.Endereco,
		}
		c.UUID = uuid.New().String()
		c.Cadastrado_em = timeNowFormatted()

		err = a.DB.CreateClient(c)
		if err != nil {
			log.Printf("Cannot save client in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		err = sendMessage(c)
		if err != nil {
			log.Printf("Falha ao enviar a mensagem err=%v \n", err)
		}

		sendResponse(w, r, &c, http.StatusOK)
	}
}

func (a *App) UpdateClientHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Cliente{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		uuid := getParameter(r)
		if uuid == "" {
			log.Printf("Cannot get uuid in request")
			sendResponse(w, r, nil, http.StatusNotFound)
			return
		}

		// Create the Client
		c := &models.Cliente{
			UUID:          uuid,
			Nome:          req.Nome,
			Endereco:      req.Endereco,
			Cadastrado_em: req.Cadastrado_em,
		}
		c.Atualizado_em = timeNowFormatted()

		// Save in DB
		err = a.DB.UpdateClient(c)
		if err != nil {
			log.Printf("Cannot update client in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, &c, http.StatusOK)
	}
}

func (a *App) GetClientsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientes, err := a.DB.GetClients()
		if err != nil {
			log.Printf("Cannot get clientes, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, &clientes, http.StatusOK)
	}
}

func (a *App) GetClientByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := getParameter(r)
		if uuid == "" {
			log.Printf("Cannot get uuid in request")
			sendResponse(w, r, nil, http.StatusNotFound)
			return
		}

		cliente, err := a.DB.GetClientById(uuid)

		if err != nil {
			log.Printf("Cannot get cliente in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, &cliente, http.StatusOK)

	}
}

func (a *App) DeleteClientHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := getParameter(r)
		fmt.Println("dados", uuid)

		if uuid == "" {
			log.Printf("Cannot get uuid in request")
			sendResponse(w, r, nil, http.StatusNotFound)
			return
		}

		// Delete in DB
		err := a.DB.DeleteClient(uuid)
		if err != nil {
			log.Printf("Cannot delete client in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, nil, http.StatusOK)
	}
}
