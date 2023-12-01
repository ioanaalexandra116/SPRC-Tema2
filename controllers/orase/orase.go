package orase

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models/database"
	"main/models/orase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrase() string {
	var selectStatement string = "SELECT * FROM orase"
	if rows, err := helpers.GetQueryResults(database.Db, selectStatement); err != nil {
		log.Println(err)
	} else {
		var tari_array = make([]orase.Oras, 0)
		for rows.Next() {
			var id int
			var id_tara int
			var nume string
			var lat float64
			var lon float64
			if err := rows.Scan(&id, &id_tara, &nume, &lat, &lon); err != nil {
				log.Println(err)
			} else {
				tara := orase.Oras{Id: id, IdTara: id_tara, Nume: nume, Lat: lat, Lon: lon}
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

func GetOraseByTara(c *gin.Context) string {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	var selectStatement string = "SELECT * FROM orase WHERE id_tara = $1"
	if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id); err != nil {
		log.Println(err)
	} else {
		var tari_array = make([]orase.Oras, 0)
		for rows.Next() {
			var id int
			var id_tara int
			var nume string
			var lat float64
			var lon float64
			if err := rows.Scan(&id, &id_tara, &nume, &lat, &lon); err != nil {
				log.Println(err)
			} else {
				tara := orase.Oras{Id: id, IdTara: id_tara, Nume: nume, Lat: lat, Lon: lon}
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

func PostOras(c *gin.Context) (int, int) {
	var oras_var orase.Oras
	if err := c.BindJSON(&oras_var); err != nil {
		log.Println(err)
		return -1, 400
	}

	var insertStatement string = "SELECT * FROM tari WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, oras_var.IdTara); err != nil {
		log.Println(err)
		return -1, 400
	} else {
		if !rows.Next() {
			return -1, 409
		}
	}

	insertStatement = "INSERT INTO orase(id_tara, nume_oras, latitudine, longitudine) VALUES($1, $2, $3, $4)"
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, oras_var.IdTara, oras_var.Nume, oras_var.Lat, oras_var.Lon); err != nil {
		log.Println(err)
	} else {
		insertStatement = "SELECT id FROM orase WHERE nume_oras = $1 AND id_tara = $2"
		if rows, err := helpers.GetQueryResults(database.Db, insertStatement, oras_var.Nume, oras_var.IdTara); err != nil {
			log.Println(err)
		} else {
			if rows.Next() {
				var id int
				if err := rows.Scan(&id); err != nil {
					log.Println(err)
				} else {
					return id, 201
				}
			}
		}
	}
	return -1, 400
}

func PutOras(c *gin.Context) int {
	var oras_var orase.Oras
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}

	var insertStatement string = "SELECT * FROM orase WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 404
		}
	}

	if err := c.BindJSON(&oras_var); err != nil {
		log.Println(err)
		return 400
	}

	insertStatement = "SELECT * FROM tari WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, oras_var.IdTara); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 409
		}
	}

	insertStatement = "UPDATE orase SET id_tara = $1, nume_oras = $2, latitudine = $3, longitudine = $4 WHERE id = $5"
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, oras_var.IdTara, oras_var.Nume, oras_var.Lat, oras_var.Lon, id); err != nil {
		log.Println(err)
		return 400
	}
	return 200
}

func DeleteOras(c *gin.Context) int {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}

	var insertStatement string = "SELECT * FROM orase WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 404
		}
	}

	insertStatement = "DELETE FROM orase WHERE id = $1"
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	}
	return 200
}