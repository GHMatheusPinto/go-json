package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Usuarios struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"-"`
	Email string `json:"email"`
}

var usuario = []Usuarios{
	Usuarios{ID: "1", Nome: "Matheus Pinto", Senha: "1234", Email: "matheus@teste.com"},
	Usuarios{ID: "2", Nome: "Rafaele Pinto", Senha: "4321", Email: "rafaele@teste.com"},
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":5000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&usuario); err != nil {
		log.Fatal(err)
	}
}
