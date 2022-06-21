package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func whatTimeIsIt(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(w, "<h1>Time: %s</h1>", s)
}

func main() {
	http.HandleFunc("/time", whatTimeIsIt)
	log.Println("Listening....")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
