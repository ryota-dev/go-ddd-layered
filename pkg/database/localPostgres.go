package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewLocalPostgres() *gorm.DB {
	// https://github.com/go-gorm/postgres
	dsn := "host=db user=postgres password=password dbname=db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("Fail to generate lcoal postgres clinet: %#v", err.Error())
	}

	return db
}
