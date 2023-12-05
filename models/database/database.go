package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)
const (
  host     = "postgres"
  port     = 5432
  user     = "admindb"
  password = "password"
  dbname   = "postgres"
  sslmode  = "require"
)

var Db *sql.DB

func Start() *sql.DB {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err := sql.Open("postgres", postgresqlDbInfo)

	if err != nil {
		panic(err)
	}

	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected to database!")
	return Db
}
