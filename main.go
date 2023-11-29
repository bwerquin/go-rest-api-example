package main

import (
	"fmt"
	"go-rest-api-example/helpers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	
	log.Println("Starting API server")
	
	//Read config file
	helpers.ReadConfig()

	//Init log system
	initLog()
	log.Println("Log configured")

	//Get Public Key from Oauth server
	helpers.InitializeOauthPublicKey()
	log.Println("Public Key configured")
	
	//Create a new router
	router := mux.NewRouter()
	log.Println("Creating routes")
	
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	
	//specific secure endpoints
	router.Handle("/secure", helpers.Protect(http.HandlerFunc(SecureZone))).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":"+helpers.AppConfig.PORT, router)
	log.Println("API ready to listen and serve")
}

func initLog() {
	//add external file for logging
	f, err := os.OpenFile(helpers.AppConfig.LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func SecureZone(w http.ResponseWriter, r *http.Request) {
	log.Println("entering secure zone end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API secure zone")
}