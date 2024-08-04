package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	fmt.Println("Running WEB ...")

	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}