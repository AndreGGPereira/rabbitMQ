package database

import (
	"api/app/models"
)

//CreateClient realiza o inserte do cliente
func (d *DB) CreateClient(c *models.Cliente) error {
	_, err := d.db.Exec(insertClientSchema, c.UUID, c.Nome, c.Endereco, c.Cadastrado_em)
	if err != nil {
		return err
	}
	return err
}

//UpdateClient realiza o update do cliente
func (d *DB) UpdateClient(c *models.Cliente) error {
	_, err := d.db.Exec(updateClientSchema, c.Nome, c.Endereco, c.Atualizado_em, c.UUID)
	if err != nil {
		return err
	}
	return err
}

//GetClientById busta uma cliente por ID
func (d *DB) GetClientById(uuid string) (*models.ClienteResponse, error) {
	var cliente models.ClienteResponse
	err := d.db.Get(&cliente, selectClientByIDSchema, uuid)
	if err != nil {
		return &cliente, err
	}
	return &cliente, err
}

//GetClients seleciona todos os clientes
func (d *DB) GetClients() ([]*models.ClienteResponse, error) {
	var clientes []*models.ClienteResponse
	err := d.db.Select(&clientes, selectClientsSchema)
	if err != nil {
		return clientes, err
	}
	return clientes, nil
}

//DeleteClient deleta cliente
func (d *DB) DeleteClient(uuid string) error {
	_, err := d.db.Exec(deleteClientSchema, uuid)
	if err != nil {
		return err
	}
	return nil
}
