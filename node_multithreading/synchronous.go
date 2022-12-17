package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start_time := time.Now()
	data, err := os.ReadFile("./uris.txt")
	if err != nil {
		log.Fatal("Could not read uri file :(")
	}
	sites := strings.Split(string(data), "\n")
	for _, uri := range sites {
		uri = strings.ReplaceAll(uri, "\r", "")
		resp, err := http.Get(uri)
		if err != nil {
			log.Fatalf("Could not get request to %v\n%v", uri, err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Something went wrong.")
		}
		filename := "./" + "/data" + uri[strings.Index(uri, ":")+2:strings.Index(uri, ".")]
		err = os.WriteFile(filename, []byte(body), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	duration := time.Since(start_time)
	fmt.Println(duration)
}
