package models

import "database/sql"

type Client struct {
	UUID       string `json:"uuid,omitempty"`
	Name       string `json:"name,omitempty"`
	Address    string `json:"address,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

type ClientResponse struct {
	UUID       string         `db:"uuid"`
	Name       sql.NullString `db:"name"`
	Address    sql.NullString `db:"address"`
	Created_at sql.NullString `db:"created_at"`
	Updated_at sql.NullString `db:"updated_at"`
}
