package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	NomClasse      string
	Filiere        string
	Niveau         string
	NombreEtudiant int
	Etudiants      []Etudiant
}

type Etudiant struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   string
}

func main() {
	Page := Page{
		NomClasse:      "B1 Informatique",
		Filiere:        "Informatique",
		Niveau:         ": Bachelor 1",
		NombreEtudiant: 30,
		Etudiants: []Etudiant{
			{Nom: "Chiotti", Prenom: "Yolan", Age: 18, Sexe: "M"},
			{Nom: "Nom", Prenom: "Prénom", Age: 21, Sexe: "F"},
		},
	}
	temp, errTemp := template.ParseGlob("promo/*.html")
	if errTemp != nil {
		fmt.Printf("Error: %v\n", errTemp)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", Page)
	})

	http.ListenAndServe("localhost:8080", nil)
}
