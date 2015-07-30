package main

import (
	"fmt"

	"github.com/mattes/migrate/migrate"
)

func RunMigrations() {
	allErrors, ok := migrate.UpSync(migrateUrl, "./migrations")
	if !ok {
		fmt.Println("Oh no ...")

		for _, err := range allErrors {
			fmt.Println(err)
		}
	}
}
