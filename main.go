package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
)

func main() {
	// Create a connection to Cassandra
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "university_of_houston"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("Error connecting to Cassandra: %v", err)
		return
	}
	defer session.Close()

	fmt.Println("Connected to Cassandra!")

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
