package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {

	log.Println("Connected to Database!", pgConnStr)
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Printf("Cannot Open DB. err=%v \n", err)
		return err
	}
	log.Println("Connected to Database!")

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
