package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/utils"
)

func main() {
	utils.LoadTemplates()
	r := router.Generate()

	fmt.Println("WEB listening in port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}