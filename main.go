package main

import (
	"log"
	"net/http"

	"app/src/landing"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// router
	r.GET("/", landing.Home)
	r.NotFound = http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":3000", r))
}
