package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":9001"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super genius stuff")
	fmt.Println("Super secret stuff")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	fmt.Println("Rest server")
	handleRequests()
}
