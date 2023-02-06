package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Reading struct {
	Name     string  `json:"name"`
	AuthCode string  `json:"auth-code"`
	Temp     float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
	Pressure float64 `json:"pressure"`
}

func write_reading(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	c.String(http.StatusOK, "Success!")
	if err != nil {
		log.Fatal("Something went wrong! :( \n")
	}
	var p Reading
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatal("Something went wrong! :( \n")
	}
	filename := "./logs/" + strings.ToLower(strings.Replace(p.Name, " ", "", -1)) + "_log.txt"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Something went wrong! :( \n")
	}
	defer file.Close() //Defer the closing of the file until the program ends
	now := time.Now().String()
	now = now[:strings.Index(now, ".")]
	text := fmt.Sprintf("%v %v %v %v\n", now, p.Temp, p.Humidity, p.Pressure)
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal("Something went wrong! :( \n")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})
	r.POST("/posttemp", write_reading)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
