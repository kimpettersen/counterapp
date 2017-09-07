package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received a request")
		w.Write([]byte("Counter accounted for"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
