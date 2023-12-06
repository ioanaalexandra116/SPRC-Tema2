package database

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
  user     = "admindb"
  password = "password"
  dbname   = "postgres"
  sslmode  = "require"
  retryInterval = time.Second * 2
  maxRetries = 10
)

var Db *sql.DB

func Start() *sql.DB {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	for i := 0; i < maxRetries; i++ {
		Db, err := sql.Open("postgres", postgresqlDbInfo)
		if err != nil {
			log.Printf("Error connecting to database (attempt %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(retryInterval)
			continue
		}

		err = Db.Ping()
		if err == nil {
			log.Println("Successfully connected to the database!")
			return Db
		}

		log.Printf("Error pinging database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	log.Fatal("Failed to connect to the database after multiple attempts.")
	return nil
}
