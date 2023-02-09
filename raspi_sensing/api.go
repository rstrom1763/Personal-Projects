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

// Struct type that holds the data from the reading
type Reading struct {
	Name     string  `json:"name"`
	AuthCode string  `json:"auth-code"`
	Temp     float64 `json:"temp"`
	Humidity float64 `json:"humidity"`
	Pressure float64 `json:"pressure"`
}

// Function that takes the sent json and records it to the log
func write_reading(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body) //Read the posted data
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, "Success!") //Send the success message

	var p Reading                  //Variable to hold the data
	err = json.Unmarshal(data, &p) //Unmarshall the data and create a struct of type Reading
	if err != nil {
		log.Fatal(err)
	}

	//Name of the log file to save the reading to
	//Based on the name of the node given in the data
	filename := "./logs/" + strings.ToLower(strings.Replace(p.Name, " ", "", -1)) + "_log.txt"

	//Open the log file for appending
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //Defer the closing of the file until the program ends

	//Get current timestamp
	//Cut off last part to make it more concise
	now := time.Now().String()
	now = now[:strings.Index(now, ".")]

	//Message to append to the log file
	text := fmt.Sprintf("%v %v %v %v\n", now, p.Temp, p.Humidity, p.Pressure)

	//Append the string to the log file
	_, err = file.WriteString(text)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := ":8081"              //Port to listen on
	gin.SetMode(gin.ReleaseMode) //Turn off debugging mode
	r := gin.Default()           //Initialize Gin

	//Route for testing functionality
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	//Post where the reading gets posted to
	r.POST("/posttemp", write_reading)

	r.Run(port)                                 //Start running the Gin server
	fmt.Printf("Listening on port %v...", port) //Notifies that server is running on X port
}
