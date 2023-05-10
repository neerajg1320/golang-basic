package main

import (
	"fmt"
	"net/http"

	"github.com/neeraj76/golang-basic/pkg/handlers"
)

const portNumber = 8080

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprint("Starting application on port", portNumber))

	_ = http.ListenAndServe(fmt.Sprint(":", portNumber), nil)
}
