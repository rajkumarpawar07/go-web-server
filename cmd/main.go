package main

import (
	"fmt"
	"net/http"
	// http package for creating http server
	"github.com/go-chi/chi"
	// to log error
	log "github.com/sirupsen/logrus"
	// import the package that contains the "handlers" module
	"go_tutorials/webserver/internal/handlers"
)

func main() {
	fmt.Print("Hello World")
	// to use logger
	log.SetReportCaller(true)
	// create a new router
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go api Service...")
	// start the server with http package
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}