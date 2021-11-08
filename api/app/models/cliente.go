package models

import "database/sql"

type Cliente struct {
	UUID          string `json:"uuid,omitempty"`
	Nome          string `json:"nome,omitempty"`
	Endereco      string `json:"endereco,omitempty"`
	Cadastrado_em string `json:"cadastrado_em,omitempty"`
	Atualizado_em string `json:"atualizado_em,omitempty"`
}

type ClienteResponse struct {
	UUID          string         `db:"uuid"`
	Nome          sql.NullString `db:"nome"`
	Endereco      sql.NullString `db:"endereco"`
	Cadastrado_em sql.NullString `db:"cadastrado_em"`
	Atualizado_em sql.NullString `db:"atualizado_em"`
}
