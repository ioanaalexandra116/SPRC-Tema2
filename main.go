package main

import (
	"main/models/database"
	"main/handlereq"
)

func main() {
    database.Db = database.Start()

    defer database.Db.Close()

    handlereq.HandleRequests()
    
}