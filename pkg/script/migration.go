package main

import (
	"fmt"
	"go-ddd-layered/infrastructure/dbmodel"
	"go-ddd-layered/pkg/database"
)

func main() {
	db := database.NewLocalPostgres()

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	if err := db.Debug().AutoMigrate(&dbmodel.User{}); err != nil {
		panic(err)
	}

	fmt.Println("migration completed")
}
