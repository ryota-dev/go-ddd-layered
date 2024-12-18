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

	for i := 0; i < 100; i++ {
		db.Create(&dbmodel.User{
			Name:  fmt.Sprintf("user-%d", i),
			Email: fmt.Sprintf("user-%d@example.com", i),
		})
	}

	fmt.Println("migration completed")
}
