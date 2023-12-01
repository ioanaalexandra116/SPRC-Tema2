package main

import (
	"main/models/database"
	"main/handlereq"
) // import "main/models/tari"

func main() {
    database.Db = database.Start()

    defer database.Db.Close()

    handlereq.HandleRequests()
    
}