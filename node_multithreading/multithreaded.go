package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup //Create the waitgroup object

func get_html(uri string) {

	defer wg.Done() //Schedule with the waitgroup
	uri = strings.ReplaceAll(uri, "\r", "")
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("Could not get request to %v:%v", uri, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Something went wrong.")
	}
	filename := "./" + "data/" + uri[strings.Index(uri, ":")+2:strings.Index(uri, ".")]
	err = os.WriteFile(filename, []byte(body), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start_time := time.Now()
	data, err := os.ReadFile("./uris.txt")
	if err != nil {
		log.Fatal("Could not read uri file :(")
	}
	sites := strings.Split(string(data), "\n")
	for _, uri := range sites {
		wg.Add(1) //Add the Go routine to the waitlist
		go get_html(uri)
	}
	wg.Wait() //Wait for all the Go routines to finish
	duration := time.Since(start_time)
	fmt.Println(duration)
}
