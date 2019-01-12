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
	// static
	fs := http.FileServer(http.Dir("./static/"))
	r.NotFound = http.StripPrefix("/static/", fs)
	// listening...
	log.Fatal(http.ListenAndServe(":3000", r))
}
