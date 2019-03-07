package main

import (
	"fmt"
	"log"
	"testProject/daemon"
	"testProject/model"
)

type Env struct {
	db model.Datastore
}

func main(){

	// Чтение yml конфигурации
	conf := daemon.Initialize()

	// Инициализация подключения БД
	db, err := model.InitDb(conf.Connect)

	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	customers, err := env.db.AllCustomers()

	for _, customer := range customers {
		fmt.Println(customer.Name)
	}
}
