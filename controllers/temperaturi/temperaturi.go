package temperaturi

import (
	"encoding/json"
	"log"
	"main/helpers"
	"main/models/database"
	"main/models/temperaturi"
	"strconv"
	"strings"
	"time"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Scanning(rows *sql.Rows) []temperaturi.TemperaturaPrint {
	var final_array = make([]temperaturi.TemperaturaPrint, 0)
	for rows.Next() {
		var id int
		var id_oras int
		var valoare float64
		var timestamp string
		if err := rows.Scan(&id, &id_oras, &valoare, &timestamp); err != nil {
			log.Println(err)
		} else {
			tara := temperaturi.TemperaturaPrint{Id: id, Valoare: valoare, Timestamp: strings.Split(timestamp, " ")[0]}
			final_array = append(final_array, tara)
		}
	}
	rows.Close()
	return final_array
}

func GetTemperaturi(c *gin.Context) string {
	var temp_array = make([]temperaturi.TemperaturaPrint, 0)
	var selectStatement string
	params := c.Request.URL.Query()
	if params["from"] == nil {
		params["from"] = append(params["from"], "")
	}
	if params["until"] == nil {
		params["until"] = append(params["until"], "")
	}
	if params["lat"] == nil {
		params["lat"] = append(params["lat"], "")
	}
	if params["lon"] == nil {
		params["lon"] = append(params["lon"], "")
	}
	from := params["from"][0]
	lat := params["lat"][0]
	lon := params["lon"][0]
	until := params["until"][0]
	if from != "" && until != "" && lat != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1 AND longitudine = $2) AND timestamp::date BETWEEN $3 AND $4"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, lon, from, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" && lat != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1 AND longitudine = $2) AND timestamp::date >= $3"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, lon, from); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" && lat != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1 AND longitudine = $2) AND timestamp::date <= $3"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, lon, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" && from != ""  && lat != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1) AND timestamp::date BETWEEN $2 AND $3"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, from, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" && from != ""  && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE longitudine = $1) AND timestamp::date BETWEEN $2 AND $3"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lon, from, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" && until != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE timestamp::date BETWEEN $1 AND $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, from, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" && lat != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1) AND timestamp::date >= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, from); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE longitudine = $1) AND timestamp::date >= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lon, from); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" && lat != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1) AND timestamp::date <= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE longitudine = $1) AND timestamp::date <= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lon, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if lat != "" && lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1 AND longitudine = $2)"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat, lon); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE timestamp::date >= $1"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, from); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE timestamp::date <= $1"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if lat != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE latitudine = $1)"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lat); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if lon != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE longitudine = $1)"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, lon); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else {
		selectStatement = "SELECT * FROM temperaturi"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	}

	if tari_json, err := json.Marshal(temp_array); err != nil {
		log.Println(err)
	} else {
		return string(tari_json)
	}

	return "[]"
}

func GetTemperaturiByOras(c *gin.Context) string {
	var selectStatement string
	var temp_array = make([]temperaturi.TemperaturaPrint, 0)
	id_oras := c.Param("id_oras")
	params := c.Request.URL.Query()
	if params["from"] == nil {
		params["from"] = append(params["from"], "")
	}
	if params["until"] == nil {
		params["until"] = append(params["until"], "")
	}
	from := params["from"][0]
	until := params["until"][0]
	if from != "" && until != "" {
        selectStatement = "SELECT * FROM temperaturi WHERE id_oras = $1 AND timestamp::date BETWEEN $2 AND $3"
        if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_oras, from, until); err != nil {
            log.Println(err)
        } else {
            temp_array = Scanning(rows)
        }
    } else if from != "" {
        selectStatement = "SELECT * FROM temperaturi WHERE id_oras = $1 AND timestamp::date >= $2"
        if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_oras, from); err != nil {
            log.Println(err)
        } else {
            temp_array = Scanning(rows)
        }
    } else if until != "" {
        selectStatement = "SELECT * FROM temperaturi WHERE id_oras = $1 AND timestamp::date <= $2"
        if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_oras, until); err != nil {
            log.Println(err)
        } else {
            temp_array = Scanning(rows)
        }
    } else {
        selectStatement = "SELECT * FROM temperaturi WHERE id_oras = $1"
        if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_oras); err != nil {
            log.Println(err)
        } else {
            temp_array = Scanning(rows)
        }
    }

	if tari_json, err := json.Marshal(temp_array); err != nil {
		log.Println(err)
	} else {
		
		return string(tari_json)
	}

    return "[]"
}

func GetTemperaturiByTara(c *gin.Context) string {
	var temp_array = make([]temperaturi.TemperaturaPrint, 0)
	var selectStatement string
	id_tara := c.Param("id_tara")
	params := c.Request.URL.Query()
	if params["from"] == nil {
		params["from"] = append(params["from"], "")
	}
	if params["until"] == nil {
		params["until"] = append(params["until"], "")
	}
	from := params["from"][0]
	until := params["until"][0]

	if from != "" && until != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE id_tara = $1) AND timestamp::date BETWEEN $2 AND $3"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_tara, from, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if from != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE id_tara = $1) AND timestamp::date >= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_tara, from); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else if until != "" {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE id_tara = $1) AND timestamp::date <= $2"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_tara, until); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	} else {
		selectStatement = "SELECT * FROM temperaturi WHERE id_oras IN (SELECT id FROM orase WHERE id_tara = $1)"
		if rows, err := helpers.GetQueryResults(database.Db, selectStatement, id_tara); err != nil {
			log.Println(err)
		} else {
			temp_array = Scanning(rows)
		}
	}

	if tari_json, err := json.Marshal(temp_array); err != nil {
		log.Println(err)
	} else {
		return string(tari_json)
	}

	return "[]"
}

func PostTemperatura(c *gin.Context) (int, int) {
	var temperatura temperaturi.Temperatura
	if err := c.BindJSON(&temperatura); err != nil {
		log.Println(err)
		return -1, 400
	}

	if temperatura.IdOras == 0 || temperatura.Valoare == nil {
		return -1, 400
	}

	var insertStatement string = "SELECT * FROM orase WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, temperatura.IdOras); err != nil {
		log.Println(err)
		return -1, 400
	} else {
		if !rows.Next() {
			return -1, 404
		}
	}

	insertStatement = "INSERT INTO temperaturi(id_oras, valoare, timestamp) VALUES($1, $2, $3)"
	timezone := time.Now().Format("2006-01-02 15:04:05.99")
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, temperatura.IdOras, temperatura.Valoare, timezone); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return -1, 409
		}
	} else {
		insertStatement = "SELECT id FROM temperaturi WHERE id_oras = $1 AND timestamp = $2"
		if rows, err := helpers.GetQueryResults(database.Db, insertStatement, temperatura.IdOras, timezone);
		err != nil {
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

func PutTemperatura(c *gin.Context) int {
	var temperatura temperaturi.Temperatura
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}

	if err := c.BindJSON(&temperatura); err != nil {
		log.Println(err)
		return 400
	}

	if temperatura.IdOras == 0 || temperatura.Valoare == nil || temperatura.Id == 0 {
		return 400
	}

	if temperatura.Id != id {
		return 4090
	}

	var insertStatement string = "SELECT * FROM temperaturi WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 4040
		}
	}

	insertStatement = "SELECT * FROM orase WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, temperatura.IdOras); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 404
		}
	}

	insertStatement = "UPDATE temperaturi SET id_oras = $2, valoare = $3 WHERE id = $4"
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, temperatura.IdOras, temperatura.Valoare, id); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 409
		}
		return 400
	}
	return 200
}

func DeleteTemperatura(c *gin.Context) int {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		return 400
	}

	var insertStatement string = "SELECT * FROM temperaturi WHERE id = $1"
	if rows, err := helpers.GetQueryResults(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	} else {
		if !rows.Next() {
			return 404
		}
	}

	insertStatement = "DELETE FROM temperaturi WHERE id = $1"
	if _, err := helpers.ExecuteStatement(database.Db, insertStatement, id); err != nil {
		log.Println(err)
		return 400
	}
	return 200
}