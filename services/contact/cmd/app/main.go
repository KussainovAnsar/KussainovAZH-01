package main

import (
	"architecture_go/pkg/store/postgres"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	host := "localhost"
	port := "5433"
	user := "postgres"
	password := "123456"
	dbname := "postgres"

	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	defer db.Close()

	fmt.Println("Connected to the database!")
}
