package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"time"

	_ "github.com/lib/pq"
)
const (
  host     = "DB_HOST"
  port     = "DB_PORT"
  user     = "DB_USER"
  password = "DB_PASSWORD"
  dbname   = "DB_NAME"
  sslmode  = "require"
  retryInterval = time.Second * 2
  maxRetries = 10
)

var Db *sql.DB

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func Start() *sql.DB {
	host := getEnvOrDefault(host, "localhost")
	port := getEnvOrDefault(port, "5432")
	user := getEnvOrDefault(user, "postgres")
	password := getEnvOrDefault(password, "postgres")
	dbname := getEnvOrDefault(dbname, "postgres")
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
