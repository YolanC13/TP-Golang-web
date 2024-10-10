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

type View struct {
	ViewCounter int
	IsOdd       bool
}

type Utilisateur struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Sexe          string
	Errors        []string
}

var utilisateurData Utilisateur

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

	View := View{
		ViewCounter: 0,
		IsOdd:       false,
	}

	temp, errTemp := template.ParseGlob("promo/*.html")
	if errTemp != nil {
		fmt.Printf("Error: %v\n", errTemp)
		return
	}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", Page)
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		View.ViewCounter++
		if View.ViewCounter%2 == 1 {
			View.IsOdd = true
		} else {
			View.IsOdd = false
		}
		temp.ExecuteTemplate(w, "change", View)
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "userForm", Page)
	})

	http.HandleFunc("/user/treatment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			utilisateurData.Errors = nil
			r.ParseForm()
			utilisateurData.Nom = r.FormValue("nom")
			utilisateurData.Prenom = r.FormValue("prenom")
			utilisateurData.DateNaissance = r.FormValue("age")
			utilisateurData.Sexe = r.FormValue("sexe")
			if len(utilisateurData.Nom) < 1 || len(utilisateurData.Nom) > 32 {
				utilisateurData.Errors = append(utilisateurData.Errors, "Le nom doit être compris entre 1 et 32 caractères")
			}
			if len(utilisateurData.Prenom) < 1 || len(utilisateurData.Prenom) > 32 {
				utilisateurData.Errors = append(utilisateurData.Errors, "Le prénom doit être compris entre 1 et 32 caractères")
			}
			if utilisateurData.Errors != nil {
				http.Redirect(w, r, "/user/error", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/user/display", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "userDisplay", utilisateurData)
	})

	http.HandleFunc("/user/error", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "userError", utilisateurData)
	})

	http.Handle("/promo/", http.StripPrefix("/promo/", http.FileServer(http.Dir("./promo"))))
	http.ListenAndServe("localhost:8080", nil)
}
