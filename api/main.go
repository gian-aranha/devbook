package main

import (
	"api/src/router"
	"fmt"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Running API...")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}