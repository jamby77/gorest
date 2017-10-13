package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	populatePeople()
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8910", router))
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index + 1])
			break
		}
	}
	GetPeople(w, r)
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person

	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]

	storePerson(person)
	json.NewEncoder(w).Encode(person)
}
func storePerson(person Person) {
	people = append(people, person)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "No person found", 404)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func populatePeople() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City 1", State: "State 1"}})
	people = append(people, Person{ID: "2", Firstname: "John 2", Lastname: "Doe 2", Address: &Address{City: "City 2", State: "State 2"}})
	people = append(people, Person{ID: "3", Firstname: "John 3", Lastname: "Doe 3", Address: &Address{City: "City 3", State: "State 3"}})
	people = append(people, Person{ID: "4", Firstname: "John 4", Lastname: "Doe 4"})
}
