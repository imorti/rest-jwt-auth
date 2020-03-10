package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("14PacCity")

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

func main() {
	fmt.Println("simple client is running")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Something went wrong GeneratingJWT: %s", err.Error())
	}

	fmt.Println(tokenString)
}
