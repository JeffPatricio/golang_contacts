package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    . "golang_rest_api/controllers"
)

func main() {
    router := mux.NewRouter()
    LoadDefaultPeople()
    router.HandleFunc("/contacts", GetPeople).Methods("GET")
    router.HandleFunc("/contacts/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contacts", CreatePerson).Methods("POST")
    router.HandleFunc("/contacts/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":4000", router))
}

