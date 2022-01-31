package database

var selectClientsSchema = `
SELECT * from Client
`
var selectClientByIDSchema = `
SELECT * from Client where uuid =$1
`

var insertClientSchema = `
INSERT INTO Client(uuid, name, address, created_at) VALUES($1,$2,$3,$4)
`

var updateClientSchema = `
UPDATE Client set name = $1, address = $2, updated_at = $3  where uuid =$4
`

var deleteClientSchema = `
DELETE from Client where uuid = $1
`
