package database

import "fmt"

var (
	dbUsername = "postgres"
	dbHost     = "localhost"
	dbPassword = "andre110407"
	dbname     = "mydatabase"
	pgConnStr  = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbname, dbUsername, dbPassword)
)

/* var (
	dbUsername = "postgres"
	dbHost     = "db"
	dbPassword = "postgres"
	dbname     = "mydatabase"
	pgConnStr  = fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		dbHost, dbname, dbUsername, dbPassword)
) */
