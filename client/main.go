package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte(os.Getenv("MY_JWT_TOKEN"))
var port = ":9000"

func homePage(w http.ResponseWriter, r *http.Request) {

	validToken, err := GenerateJWT()

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)

	// fmt.Fprintf(w, "Hello Fantastic. What's going on? ")
	// fmt.Println("Endpoint Hit: homePage")

}

// GenerateJWT - generates JWT token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "imorti"
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Error generating JWT: %s", err.Error())
		return "", err
	}

	return tokenString, err

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {

	handleRequests()
}
