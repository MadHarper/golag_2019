package model

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type Datastore interface {
	AllOrders() ([]*Order, error)
	AllCustomers() ([]*Customer, error)
}

func InitDb(conection string) (*DB, error) {

	db, err := sql.Open("postgres", conection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

