package controllers

import (
    "strconv"
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    . "golang_rest_api/model"
)

var people []Person
var lastId int

func LoadDefaultPeople(){
    people = append(people, Person{ ID: "1", Firstname: "Michael", Lastname: "Jack", Address: &Address{ City: "City Good", State: "State of Wolves" }})
    people = append(people, Person{ ID: "2", Firstname: "Ayk", Lastname: "Mauriss", Address: &Address{ City: "Zaniny", State: "State Year" }})
    people = append(people, Person{ ID: "3", Firstname: "Franx", Lastname: "Suneda" })
    lastId = 3
}

func GetPeople(res http.ResponseWriter, req *http.Request) {
    json.NewEncoder(res).Encode(people)
}

func GetPerson(res http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(res).Encode(item)
            return
        }
    }
    json.NewEncoder(res).Encode(&Person{})
}

func CreatePerson(res http.ResponseWriter, req *http.Request) {
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    lastId = lastId + 1
    person.ID = strconv.Itoa(lastId)
    people = append(people, person)
    json.NewEncoder(res).Encode(person)
}

func DeletePerson(res http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            json.NewEncoder(res).Encode(item)
            return 
        }
    }
    json.NewEncoder(res).Encode(&Person{})
}