package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"/src/github.com/andreggpereira/nuveo/api/database"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const layoutDate = "2006-01-02T15:04:05.999999"

type Cliente struct {
	UUID          string `json:"uuid,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Endereco      string `json:"endereco,omitempty"`
	Cadastrado_em string `json:"cadastrado_em,omitempty"`
	Atualizado_em string `json:"atualizado_em,omitempty"`
}

//CreateCliente cadastro
func CreateCliente(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	c := &Cliente{}
	err := decoder.Decode(c)
	if err != nil {
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
		return
	}
	c.UUID = uuid.New().String()
	c.Cadastrado_em = time.Now().Format(layoutDate)

	err = createCliente(c)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	err1 := sendMessage(c)
	if err1 != nil {
		json.NewEncoder(w).Encode("Error created message")
		fmt.Printf("Error: %s", err)
	}

	json.NewEncoder(w).Encode("OK")
}

func createCliente(c *Cliente) error {

	stmt, err := database.DB.Prepare("insert into Cliente(uuid, Nome, Endereco) values($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.UUID, c.Nome, c.Endereco)
	if err != nil {
		return err
	}
	return nil
}

//UpdateCliente cadastro
func UpdateCliente(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	c := &Cliente{}
	err := decoder.Decode(c)
	if err != nil {
		json.NewEncoder(w).Encode("Não foi possível realizar o cadastro")
		return
	}

	c.Atualizado_em = time.Now().Format(layoutDate)
	err = updateCliente(c)
	if err != nil {
		json.NewEncoder(w).Encode("Failed Cliente creation")
		return
	}

	//adicionar a fila de mensageria
	//json.NewEncoder(w).Encode(Message{Message: "Cliente created!!", Status: 200})
}

func updateCliente(c *Cliente) error {
	/* stmt, err := database.DB.Prepare("update Cliente set nome = $1, endereco = $2, atualizado_em = $3  where uuid =$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	c.Atualizado_em = time.Now().String()

	//c.Atualizado_em = time.Now().Format("02/01/2006 03:04:05")
	//2006-01-02T15:04:05-0700"
	c.Atualizado_em = time.Now().Format("2006-01-02T15:04:05-0700")

	fmt.Println(c.Atualizado_em)
	_, err = stmt.Exec(c.Nome, c.Endereco, c.Atualizado_em, c.UUID)
	if err != nil {
		fmt.Println("Erro", err)
		return err
	} */
	return nil
}

func DeleteCliente(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uiid, ok := vars["uuid"]

	if uiid == "" || !ok {
		json.NewEncoder(w).Encode("Failed decode")
		return
	}

	err := deleteCliente(uiid)
	if err != nil {
		//	json.NewEncoder(w).Encode(Message{Message: "Failed deleted", Status: 400})
		return
	}
	//json.NewEncoder(w).Encode(Message{Message: "Request completed successfully", Status: 200})
}

func deleteCliente(uiid string) error {
	/*
		stmt, err := database.DB.Prepare("delete from cliente where uuid = $1")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(uiid)
		if err != nil {
			return err
		}
	*/
	return nil
}

func GetCliente(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	uiid, ok := vars["uuid"]

	if uiid == "" || !ok {
		json.NewEncoder(w).Encode("Failed decode")
		return
	}

	c, err := getClienteID(uiid)

	if err != nil {
		//	json.NewEncoder(w).Encode(Message{Message: "Failed get Permission", Status: 400})
		return
	}
	json.NewEncoder(w).Encode(&c)
}

func getClienteID(uuid string) (*Cliente, error) {
	c := Cliente{}

	/*
		 rows, err := database.DB.Query("select * from cliente where uuid = $1", uuid)

		if err != nil {
			return nil, err
		}
		for rows.Next() {
			rows.Scan(&c.UUID, &c.Nome, &c.Endereco, &c.Cadastrado_em, &c.Atualizado_em)
			c = Cliente{UUID: c.UUID, Nome: c.Nome, Endereco: c.Endereco, Cadastrado_em: c.Cadastrado_em, Atualizado_em: c.Atualizado_em}
		}
		database.DB.Close()
		if err != nil {
			return &c, err
		} */
	return &c, nil
}

func GetClientes(w http.ResponseWriter, req *http.Request) {

	cs, err := getClientes()
	if err != nil {
		json.NewEncoder(w).Encode("Failed list Clientes")
	}
	json.NewEncoder(w).Encode(&cs)
}

func getClientes() ([]Cliente, error) {
	cs := make([]Cliente, 0)

	/*
		rows, err := database.DB.Query("SELECT uuid, nome, endereco, cadastrado_em, atualizado_em FROM Cliente")
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			c := Cliente{}
			rows.Scan(&c.UUID, &c.Nome, &c.Endereco, &c.Cadastrado_em, &c.Atualizado_em)
			cs = append(cs, Cliente{UUID: c.UUID, Nome: c.Nome, Endereco: c.Endereco, Cadastrado_em: c.Cadastrado_em, Atualizado_em: c.Atualizado_em})
		}
	*/
	return cs, nil

}
