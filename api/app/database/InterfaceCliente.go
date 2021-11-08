package database

import "api/app/models"

type InterfaceClientDB interface {
	Open() error
	Close() error
	CreateClient(c *models.Cliente) error
	UpdateClient(c *models.Cliente) error
	GetClientById(uuid string) (*models.ClienteResponse, error)
	GetClients() ([]*models.ClienteResponse, error)
	DeleteClient(uuid string) error
}
