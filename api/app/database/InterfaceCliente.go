package database

import "api/app/models"

type InterfaceClientDB interface {
	Open() error
	Close() error
	CreateClient(c *models.Client) error
	UpdateClient(c *models.Client) error
	GetClientById(uuid string) (*models.ClientResponse, error)
	GetClients() ([]*models.ClientResponse, error)
	DeleteClient(uuid string) error
}
