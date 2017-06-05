package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"{{cookiecutter.repo}}/pkg/env"
	"{{cookiecutter.repo}}/pkg/handlers"
)

var ()

func init() {
	err := env.SetUp()
	if err != nil {
		// error out with system logger 
		// as unable to parse configuration
		log.Fatal("Unable to parse configuration")
	}
}

func main() {

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	v1.Handle("/", handlers.GETBase()).Methods("GET")

	http.ListenAndServe(fmt.Sprintf(":%d", env.ListenerPort), r)
}
