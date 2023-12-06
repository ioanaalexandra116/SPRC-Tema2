package tari

import (
	"encoding/json"	
	"log"
	"main/helpers"
	"main/models/database"
	"main/models/tari"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTari() string {
	var selectStatement string = "SELECT * FROM tari"

	if rows, err := helpers.GetQueryResults(database.Db, selectStatement); err != nil {
		log.Println(err)
	} else {
		var tari_array = make([]tari.Tara, 0)
		for rows.Next() {
			var id int
			var nume string
			var lat *float64
			var lon *float64
			if err := rows.Scan(&id, &nume, &lat, &lon); err != nil {
				log.Println(err)
			} else {
				tara := tari.Tara{Id: id, Nume: nume, Lat: lat, Lon: lon}
				tari_array = append(tari_array, tara)
			}
		}
		if tari_json, err := json.Marshal(tari_array); err != nil {
			log.Println(err)
		} else {
			return string(tari_json)
		}
	}
	return "[]"
}

func PostTara(c *gin.Context) int {
	var tara_var tari.Tara
	if err := c.BindJSON(&tara_var); err != nil {
		log.Println(err)
		return 400
	}

	if tara_var.Nume == "" || tara_var.Lat == nil || tara_var.Lon == nil {
		return 400
	}

	var insertStatement string = "INSERT INTO tari(nume_tara, latitudine, longitudine) VALUES($1, $2, $3)"

	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, tara_var.Nume, tara_var.Lat, tara_var.Lon); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 409
		}
	} else {
		insertStatement = "SELECT id FROM tari WHERE nume_tara = $1"
		if rows, err := helpers.GetQueryResults(database.Db, insertStatement, tara_var.Nume); err != nil {
			log.Println(err)
		} else {
			if rows.Next() {
				var id int
				if err := rows.Scan(&id); err != nil {
					log.Println(err)
				} else {
					return id
				}
			}
		}
	}
	return 400
}

func PutTara(c *gin.Context) int {
	var tara_var tari.Tara
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}
	if err := c.BindJSON(&tara_var); err != nil {
		log.Println(err)
		return 400
	}

	if tara_var.Nume == "" || tara_var.Lat == nil || tara_var.Lon == nil || tara_var.Id == 0 {
		return 400
	}

	var selectStatement string = "SELECT * FROM tari WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id); err != nil {
		log.Println(err)
	} else {
		if !rows.Next() {
			return 404
		}
	}

	var updateStatement string = "UPDATE tari SET nume_tara = $1, latitudine = $2, longitudine = $3 WHERE id = $4"

	if _, err := helpers.ExecuteStatement(database.Db, updateStatement, tara_var.Nume, tara_var.Lat, tara_var.Lon, id); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") || strings.Contains(err.Error(), "violates foreign key constraint") {
			return 409
		}
	} else {
		return 200
	}
	return 400
}

func DeleteTara(c *gin.Context) int {
	idStr := c.Param("id")
	idStr = strings.TrimPrefix(idStr, ":")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}

	var selectStatement string = "SELECT * FROM tari WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id); err != nil {
		log.Println(err)
	} else {
		if !rows.Next() {
			return 404
		}
	}

	var deleteStatement string = "DELETE FROM tari WHERE id = $1"
	if _, err := helpers.ExecuteStatement(database.Db, deleteStatement, id); err != nil {
		log.Println(err)
	} else {
		return 200
	}
	return 400
}

