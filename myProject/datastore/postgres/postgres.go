package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var Db *sql.DB

func init() { //to setup the database together with code
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		// log.Printf(err)
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
