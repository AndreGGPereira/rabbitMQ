//go:build mage
// +build mage

package main

import (
	"fmt"

	"context"

	"github.com/gosidekick/migration/v3"
	_ "github.com/lib/pq"
)

// https://magefile.org

func Migration(action string) {
	source := "./app/database/migrations/"
	url := "postgres://postgres:postgres@localhost:15432/nuveo?sslmode=disable"
	_, _, err := migration.Run(context.Background(), source, url, action)
	if err != nil {
		fmt.Println("Migration err :", err)
	}
}
