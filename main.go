package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")

	data := []Status{}
	r.GET("/index", func(c *gin.Context) {
		water := rand.Intn(100)
		wind := rand.Intn(100)
		status := "AMAN"

		if water > 8 || wind > 15 {
			status = "BAHAYA"
		} else if water >= 6 || wind >= 7 {
			status = "SIAGA"
		}

		newData := Status{
			Water:  water,
			Wind:   wind,
			Status: status,
		}

		data =  append([]Status{newData}, data...)

		file, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile("test.json", file, 0644)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Air vs Wind",
			"water":   water,
			"wind":    wind,
			"status":  status,
			"history": data,
		})
	})

	r.Run(":8080")
}