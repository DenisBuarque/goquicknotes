package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func NoteList(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/layoutBase.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Error ao carregar página.", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "layoutBase", nil)
}

func NoteView(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"views/templates/layoutBase.html",
		"views/templates/pages/show.html",
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id não encontrado", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao carregar página show", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "layoutBase", id)
}

func NoteCreate(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/layoutBase.html",
		"views/templates/pages/create.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Error ao carregar página.", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "layoutBase", nil)
}

func NoteStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost) // informa o tipo de metodo obrogatório da requisição

		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido")
		//http.Error(w, "Métpdo não permidito", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma noto.")
}
