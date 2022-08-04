package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/req", func(w http.ResponseWriter, r *http.Request) {

		resp, err := http.Get("http://localhost:9000/info")
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
			return
		}
		defer resp.Body.Close()

		byteArry, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
			return
		}

		fmt.Fprintf(w, "Response: %s", string(byteArry))
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, "Error %s", err)
			return
		}

		fmt.Fprintf(w, "Hello, %s", hostname)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
