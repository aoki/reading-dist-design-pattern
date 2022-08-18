package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sleep := time.Duration((rand.Intn(5)+1)*(rand.Intn(5)+1)) * time.Second
	log.Printf("Wait: %v\n", sleep)
	time.Sleep(sleep)

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ready")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
			return
		}

		fmt.Fprintf(w, "Hello, %s %v", hostname, sleep)
	})

	log.Printf("Listen :8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
