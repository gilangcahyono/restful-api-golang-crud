package main

import (
	"fmt"
	"log"
	"net/http"
	"restfull-api-golang/src"
)

func main() {
	router := src.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Server has been running at http://localhost:8080")
}
