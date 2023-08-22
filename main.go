package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	initLog()
	log.Println("This is a test log entry")

	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}

func initLog() {
	//add external file for logging
	f, err := os.OpenFile("log/testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

func Persons(w http.ResponseWriter, r *http.Request) {
	log.Println("entering persons end point")
	var response Response
	persons := prepareResponse()

	response.Persons = persons

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func prepareResponse() []Person {
	var persons []Person

	var person Person
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "N"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "E"
	persons = append(persons, person)

	person.Id = 3
	person.FirstName = "Thomas"
	person.LastName = "E"
	persons = append(persons, person)
	return persons
}
