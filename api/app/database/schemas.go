package database

var selectClientsSchema = `
SELECT * from Cliente
`
var selectClientByIDSchema = `
SELECT * from Cliente where uuid =$1
`

var insertClientSchema = `
INSERT INTO Cliente(uuid, nome, endereco, cadastrado_em) VALUES($1,$2,$3,$4)
`

var updateClientSchema = `
UPDATE Cliente set nome = $1, endereco = $2, atualizado_em = $3  where uuid =$4
`

var deleteClientSchema = `
DELETE from Cliente where uuid = $1
`
