package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var appName string

// Listen and serve HTTP requests on port 8080
func main() {
	name, ok := os.LookupEnv("APP_NAME")
	if !ok {
		log.Fatal("missing APP_NAME env variable")
	}
	appName = name
	http.HandleFunc("/", HelloServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Greeting")

	_, err := fmt.Fprintf(w, "Hello world from %s!", appName)
	if err != nil {
		w.WriteHeader(500)
	}
}
