package handlereq

import (
	"encoding/json"
	"log"
	"main/controllers/tari"
	"main/controllers/orase"
	"main/controllers/temperaturi"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/api/countries", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(tari.GetTari()))
	})

	router.POST("/api/countries", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res := tari.PostTara(c)
			if res != -1 {
				c.JSON(201, gin.H{"id": res})
			} else {
				c.JSON(400, "Bad request")
			}
		} else {
			c.JSON(409, "Bad request")
		}
	})

	router.PUT("/api/countries/:id", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res := tari.PutTara(c)
			if res == 200 {
				c.JSON(res, "Country updated")
			} else {
				c.JSON(res, "Bad request")
			}
		} else {
			c.JSON(400, "Bad request")
		}
	})

	router.DELETE("/api/countries/:id", func(c *gin.Context) {
		res := tari.DeleteTara(c)
		if res == 200 {
			c.JSON(res, "Country deleted")
		} else {
			c.JSON(res, "Bad request")
		}
	})

	router.GET("/api/cities", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(orase.GetOrase()))
	})

	router.GET("/api/cities/country/:id", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(orase.GetOraseByTara(c)))
	})

	router.POST("/api/cities", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res, code := orase.PostOras(c)
			if res != -1 {
				c.JSON(code, gin.H{"id": res})
			} else {
				c.JSON(code, "Bad request")
			}
		} else {
			c.JSON(400, "Bad request")
		}
	})

	router.PUT("/api/cities/:id", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res := orase.PutOras(c)
			if res == 200 {
				c.JSON(res, "City updated")
			} else {
				c.JSON(res, "Bad request")
			}	
		} else {
			c.JSON(400, "Bad request")
		}
	})

	router.DELETE("/api/cities/:id", func(c *gin.Context) {
		res := orase.DeleteOras(c)
		if res == 200 {
			c.JSON(res, "City deleted")
		} else {
			c.JSON(res, "Bad request")
		}
	})

	router.GET("/api/temperatures", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(temperaturi.GetTemperaturi(c)))
	})

	router.GET("/api/temperatures/cities/:id_oras", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(temperaturi.GetTemperaturiByOras(c)))
	})

	router.GET("/api/temperatures/countries/:id_tara", func(c *gin.Context) {
		c.JSON(200, json.RawMessage(temperaturi.GetTemperaturiByTara(c)))
	})

	router.POST("/api/temperatures", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res, code := temperaturi.PostTemperatura(c)
			if res != -1 {
				c.JSON(code, gin.H{"id": res})
			} else {
				c.JSON(code, "Bad request")
			}
		} else {
			c.JSON(400, "Bad request")
		}
	})

	router.PUT("/api/temperatures/:id", func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") == "application/json" {
			res := temperaturi.PutTemperatura(c)
			if res == 200 {
				c.JSON(res, "Temperature updated")
			} else {
				c.JSON(res, "Bad request")
			}
		} else {
			c.JSON(400, "Bad request")
		}
	})

	router.DELETE("/api/temperatures/:id", func(c *gin.Context) {
		res := temperaturi.DeleteTemperatura(c)
		if res == 200 {
			c.JSON(res, "Temperature deleted")
		} else {
			c.JSON(res, "Bad request")
		}
	})

	router.Run("localhost:6000")
	log.Println("Server started on: http://localhost:8080")
}