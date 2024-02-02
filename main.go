package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "my_user"
	password = "my_password"
	dbname   = "my_database"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	for {
		log.Println("OK")
		time.Sleep(time.Second)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
