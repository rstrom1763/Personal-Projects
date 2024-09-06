package main

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// Struct type that holds the data from the reading
type Reading struct {
	Name     string  `json:"name"`
	AuthCode string  `json:"auth-code"`
	Temp     float32 `json:"temp"`
	Humidity float32 `json:"humidity"`
	Pressure float32 `json:"pressure"`
	Time     int32   `json:"time"`
}

type Temp struct {
	Temp float32 `json:"temp"`
	Time int32   `json:"time"`
}

type Hist struct {
	Temps []float32
	Times []int32
}

// Function to get the most recent 480 temperature readings
func getRecentTemperatureReadings(session *gocql.Session) (string, error) {

	query := `SELECT temp, time FROM temps WHERE name = 'office' ORDER BY time DESC LIMIT 480`

	// Execute the query and create an iterator
	iter := session.Query(query).Iter()

	// Defer close Scylla query iterator
	defer func(iter *gocql.Iter) {
		err := iter.Close()
		if err != nil {
			log.Printf("Error closing Scylla query iter: %s", err)
		}
	}(iter)

	// Prepare a slice to hold the results
	var readings []Temp
	var temp float32
	var time int32

	// Iterate over the results
	for iter.Scan(&temp, &time) {
		reading := Temp{
			Temp: temp,
			Time: time,
		}
		readings = append(readings, reading)
	}

	// Check for errors during iteration
	if err := iter.Close(); err != nil {
		return "", err
	}

	var hist Hist
	for _, reading := range readings {
		hist.Temps = append(hist.Temps, reading.Temp)
		hist.Times = append(hist.Times, reading.Time)
	}

	hist = getMinuteAverages(hist)

	histJSON, _ := json.Marshal(hist)

	return string(histJSON), nil
}

// Function that takes the received JSON and records it to the log
func insertReading(c *gin.Context, session *gocql.Session) {
	data, err := io.ReadAll(c.Request.Body) // Read the posted data
	if err != nil {
		log.Println("Error reading request body:", err)
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}

	var p Reading                  // Variable to hold the data
	err = json.Unmarshal(data, &p) // Unmarshal the data and create a struct of type Reading
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		c.String(http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Get current timestamp in unix epoch format
	now := time.Now()
	p.Time = int32(now.Unix())

	// Define the insert query
	insertQuery := `INSERT INTO temps (name, id, "auth-code", humidity, pressure, temp, time) VALUES (?, ?, ?, ?, ?, ?, ?)`

	id := gocql.TimeUUID()

	// Execute the insert query
	err = session.Query(insertQuery, p.Name, id, p.AuthCode, p.Humidity, p.Pressure, p.Temp, p.Time).Exec()
	if err != nil {
		log.Println("Failed to execute query:", err)
		c.String(http.StatusInternalServerError, "Failed to insert data")
		return
	}

	c.String(http.StatusOK, "Success!") // Send the success message
}

func getLatestReading(session *gocql.Session) (float32, error) {
	getQuery := `SELECT temp FROM temps WHERE name = 'office' ORDER BY time DESC LIMIT 1 `
	var temp Temp
	err := session.Query(getQuery).Scan(&temp.Temp)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return 0.0, nil
	}
	return temp.Temp, nil
}

func abortWithError(statusCode int, err error, c *gin.Context) {
	_ = c.AbortWithError(statusCode, err)
	c.JSON(statusCode, gin.H{"status": fmt.Sprint(err)})
}

// gzipBytes compresses the input bytes and returns the gzipped bytes.
func gzipBytes(data []byte) ([]byte, error) {
	// Create a buffer to hold the gzipped data.
	var buf bytes.Buffer

	// Create a new gzip writer with the buffer.
	gz := gzip.NewWriter(&buf)

	// Write the data to the gzip writer.
	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	// Close the gzip writer to flush any remaining data.
	err = gz.Close()
	if err != nil {
		return nil, err
	}

	// Return the gzipped data.
	return buf.Bytes(), nil
}

func generateSSL() {

	// Generate a private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Error generating private key:", err)
		return
	}

	// Generate a self-signed certificate
	certTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Fatal("Error creating certificate:", err)
		return
	}

	// Write the private key and certificate to files
	keyOut, err := os.Create("./private.key")
	if err != nil {
		log.Fatal("Error creating private key file:", err)
		return
	}

	defer func(keyOut *os.File) { // A defer for closing out the private key
		err := keyOut.Close()
		if err != nil {
			log.Fatal("Error closing private key file:", err)
		}
	}(keyOut)

	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		log.Fatal("Error creating certificate file: ", err)
		return
	}

	certOut, err := os.Create("./cert.pem")
	if err != nil {
		log.Fatal("Error creating certificate file: ", err)
		return
	}
	defer func(certOut *os.File) { // A defer for closing out the cert file
		err := certOut.Close()
		if err != nil {
			log.Fatal("Error closing certificate file: ", err)
		}
	}(certOut)

	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		log.Fatal("Error creating certificate file: ", err)
		return
	}

	fmt.Println("TLS certificate and private key generated successfully.")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // File exists
	}
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return false // Error occurred (e.g., permission denied)
}

func avg(nums []float32) float32 {
	var final float32
	for _, num := range nums {
		final += num
	}
	return final / float32(len(nums))
}

func getMinuteAverages(hist Hist) Hist {

	var newHist Hist
	var slice []float32

	for i, reading := range hist.Temps {
		slice = append(slice, reading)
		if (i+1)%4 == 0 {
			newHist.Temps = append(newHist.Temps, avg(slice))
			slice = slice[:0]
		}
	}

	for i, timeStamp := range hist.Times {
		if (i+1)%4 == 0 {
			newHist.Times = append(newHist.Times, timeStamp)
		}
	}

	return newHist
}

func main() {
	port := "8081"               // Port to listen on
	gin.SetMode(gin.ReleaseMode) // Turn off debugging mode
	r := gin.Default()           // Initialize Gin
	protocol := "https"

	// Define the ScyllaDB cluster configuration
	cluster := gocql.NewCluster("10.0.0.234") // Replace with your ScyllaDB node addresses
	cluster.Keyspace = "raspi_sensing"        // Replace with your keyspace name
	cluster.Consistency = gocql.Quorum
	chartFileBytes, _ := os.ReadFile("./chart.js")
	indexFileBytes, _ := os.ReadFile("./index.html")
	gzChartFileBytes, _ := gzipBytes(chartFileBytes)
	gzIndexFileBytes, _ := gzipBytes(indexFileBytes)

	//Generate TLS keys if they do not already exist
	if !(fileExists("./cert.pem") && fileExists("./private.key")) && protocol == "https" {
		generateSSL()
	}

	// Create a session to the cluster
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to ScyllaDB cluster: %v", err)
	}

	// Close session at the end
	defer session.Close()

	fmt.Println("Connected to ScyllaDB")

	// Route for testing reachability
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Gets the temp history html
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Encoding", "gzip")
		c.Data(http.StatusOK, "text/html", gzIndexFileBytes)
	})

	// Get the latest temperature reading
	r.GET("/temp", func(c *gin.Context) {
		temp, err := getLatestReading(session)
		if err != nil {
			abortWithError(http.StatusInternalServerError, err, c)
			return
		}
		c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprint(temp)))
	})

	// Returns the chart.js file as gzipped bytes
	r.GET("/chart.js", func(c *gin.Context) {
		c.Header("Content-Encoding", "gzip")
		c.Data(http.StatusOK, "application/javascript", gzChartFileBytes)
	})

	// Route where the reading gets posted to
	r.POST("/posttemp", func(c *gin.Context) {
		insertReading(c, session)
	})

	// Get the temperatures from the last 2 hours
	r.GET("/getHist", func(c *gin.Context) {
		temps, err := getRecentTemperatureReadings(session)
		if err != nil {
			fmt.Printf("could not get temp history: %v", err)
		}
		tempsBytes, _ := gzipBytes([]byte(temps))
		c.Header("Content-Encoding", "gzip")
		c.Data(http.StatusOK, "text/plain", tempsBytes)
	})

	fmt.Printf("Listening for %v on port %v...\n", protocol, port) //Notifies that server is running on X port
	if protocol == "http" {                                        //Start running the Gin server
		err = r.Run(":" + port)
		if err != nil {
			fmt.Println(err)
		}
	} else if protocol == "https" {
		err = r.RunTLS(":"+port, "./cert.pem", "./private.key")
		if err != nil {
			fmt.Println(err)
		}
	} else {
		log.Fatal("Something went wrong starting the Gin server")
	}
}
