package database

import (
	"api/app/models"
)

//CreateClient realiza o inserte do Client
func (d *DB) CreateClient(c *models.Client) error {
	_, err := d.db.Exec(insertClientSchema, c.UUID, c.Name, c.Address, c.Created_at)
	if err != nil {
		return err
	}
	return err
}

//UpdateClient realiza o update do Client
func (d *DB) UpdateClient(c *models.Client) error {
	_, err := d.db.Exec(updateClientSchema, c.Name, c.Address, c.Updated_at, c.UUID)
	if err != nil {
		return err
	}
	return err
}

//GetClientById busta uma Client por ID
func (d *DB) GetClientById(uuid string) (*models.ClientResponse, error) {
	var Client models.ClientResponse
	err := d.db.Get(&Client, selectClientByIDSchema, uuid)
	if err != nil {
		return &Client, err
	}
	return &Client, err
}

//GetClients seleciona todos os Clients
func (d *DB) GetClients() ([]*models.ClientResponse, error) {
	var Clients []*models.ClientResponse
	err := d.db.Select(&Clients, selectClientsSchema)
	if err != nil {
		return Clients, err
	}
	return Clients, nil
}

//DeleteClient deleta Client
func (d *DB) DeleteClient(uuid string) error {
	_, err := d.db.Exec(deleteClientSchema, uuid)
	if err != nil {
		return err
	}
	return nil
}
