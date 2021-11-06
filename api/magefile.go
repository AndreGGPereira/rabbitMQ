package main

import (
	"fmt"
	"os"
	"time"

	"context"

	"github.com/gosidekick/migration/v3"
	_ "github.com/lib/pq"
)

func NewMigration(name string) {

	t := time.Now()
	timestamp := t.Format("20060102150405")
	timestamp = timestamp + "_"
	f1, err := os.Create("/migration/" + timestamp + name + ".up.sql")

	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	f2, err := os.Create("/migration/" + timestamp + name + ".down.sql")
	if err != nil {
		fmt.Println(err)
	}
	defer f2.Close()

}

func Migration(action string) {

	url := "postgres://postgres:andre110407@localhost:5432/nuveo?sslmode=disable"
	source := "/api/migration/"
	_, _, err := migration.Run(context.Background(), source, url, action)
	if err != nil {
		fmt.Println("Migration err :", err)
	}
}
