package main

import (
	"ambedo-ws/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// TEMP ---- REMOVE
	fmt.Println("AMBEDO API IS RUNNING...")

	router := router.Generate()
	log.Fatal(http.ListenAndServe(":8080", router))
}
