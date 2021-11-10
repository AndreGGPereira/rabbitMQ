package database

import (
	"api/app/models"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	d := &DB{}
	d.Open()

	c := models.Cliente{
		UUID:     "0ddc782b-6618-4788-a65d-2499028487e6",
		Nome:     "Nome",
		Endereco: "Endereco",
	}
	timer := time.Now().Format("2006-01-02T15:04:05.999999")
	c.Cadastrado_em = timer
	_, err := d.db.Exec(insertClientSchema, c.UUID, c.Nome, c.Endereco, c.Cadastrado_em)
	assert.NoError(t, err)
}

func TestUpdateClientClient(t *testing.T) {
	d := &DB{}
	d.Open()

	c := models.Cliente{
		UUID:     "0ddc782b-6618-4788-a65d-2499028487e6",
		Nome:     "Nome Up",
		Endereco: "Endereco Up",
	}
	timer := time.Now().Format("2006-01-02T15:04:05.999999")
	c.Atualizado_em = timer
	_, err := d.db.Exec(updateClientSchema, c.Nome, c.Endereco, c.Atualizado_em, c.UUID)
	assert.NoError(t, err)
}

//GetClientById busta uma cliente por ID
func TestGetClientById(t *testing.T) {
	d := &DB{}
	d.Open()
	var cliente models.ClienteResponse
	err := d.db.Get(&cliente, selectClientByIDSchema, "0ddc782b-6618-4788-a65d-2499028487e6")
	assert.NoError(t, err)
}

//GetClients seleciona todos os clientes
func TestGetClients(t *testing.T) {
	d := &DB{}
	d.Open()
	var clientes []*models.ClienteResponse
	err := d.db.Select(&clientes, selectClientsSchema)
	assert.NoError(t, err)
}

//DeleteClient deleta cliente
func TestDeleteClient(t *testing.T) {
	d := &DB{}
	d.Open()
	_, err := d.db.Exec(deleteClientSchema, "0ddc782b-6618-4788-a65d-2499028487e6")
	assert.NoError(t, err)
}
