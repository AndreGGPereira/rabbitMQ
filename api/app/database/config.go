package database

import "fmt"

var (
	dbUsername = "postgres"
	dbHost     = "db"
	dbPassword = "postgres"
	dbname     = "nuveo"
	pgConnStr  = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbname, dbUsername, dbPassword)
)
