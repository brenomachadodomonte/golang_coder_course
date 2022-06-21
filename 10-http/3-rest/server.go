package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", UserHandler)
	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
