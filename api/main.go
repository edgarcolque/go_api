package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "https://github.com/edgarcolque/go_api/api/libs"
)

func main() {
    var router = mux.NewRouter()
    router.HandleFunc("/healthcheck", healthCheck).Methods("GET")
    router.HandleFunc("/items", handleQryMessage).Methods("GET")
    router.HandleFunc("/items/{msg}", handleUrlMessage).Methods("GET")

    fmt.Println("Running server!")
    log.Fatal(http.ListenAndServe(":3000", router))
}

func handleQryMessage(w http.ResponseWriter, r *http.Request) {
    vars := r.URL.Query()
    message := vars.Get("msg")

    json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    message := vars["msg"]

    json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode("Still alive!")
}
