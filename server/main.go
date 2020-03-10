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
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(port, nil))
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func main() {
	fmt.Println("Rest server")
	handleRequests()
}
