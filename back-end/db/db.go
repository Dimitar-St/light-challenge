package db

import (
	"light-challenge/models"
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func InitDatabase() (*pg.DB, error) {
	log.Print("Initializing database connection")
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "12345",
	})

	models := []interface{}{
		(*models.Invoice)(nil),
		(*models.Employee)(nil),
		(*models.DepartmentType)(nil),
	}

	for _, v := range models {
		err := db.Model(v).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})

		if err != nil {
			return &pg.DB{}, err
		}
	}

	log.Print("Database connection successful")

	return db, nil

}
