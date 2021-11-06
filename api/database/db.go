package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//DB teste
var DB *sqlx.DB

func init() {
	var err error

	/* 	dbUsername := os.Getenv("DB_USERNAME")
	   	dbPassword := os.Getenv("DB_PASSWORD")
	   	dbHost := os.Getenv("DB_HOST")
	   	dbTable := os.Getenv("DB_TABLE")
	   	dbPort := os.Getenv("DB_PORT")
	   	dbSSLMode := os.Getenv("DB_SSL_MODE") */
	dbUsername := "postgres"
	dbPassword := "postgres"
	dbHost := "localhost"
	dbTable := "nuveo"
	dbPort := "5432"
	dbSSLMode := "disable"

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	DB, err = sqlx.Connect("postgres", connectionString)
	fmt.Println("Testanto a conexao do banco de dados")
	if err != nil {
		fmt.Println(err)
	}
	if err = DB.Ping(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Conexão Estabelecida")
}
