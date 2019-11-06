package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.HandlerFunc("POST", "/smarthome", wrapHandler(handleSmartHome()))

	log.Fatal(http.ListenAndServe(":8081", router))
}
