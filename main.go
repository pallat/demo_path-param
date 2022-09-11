package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos/{id}", todoHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	resp := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("you need todo id %s", id),
	}

	json.NewEncoder(w).Encode(&resp)
}
