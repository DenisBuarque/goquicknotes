package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
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

func noteView(w http.ResponseWriter, r *http.Request) {

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

func noteCreate(w http.ResponseWriter, r *http.Request) {
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

func noteStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost) // informa o tipo de metodo obrogatório da requisição

		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido")
		//http.Error(w, "Métpdo não permidito", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma noto.")
}

func main() {
	fmt.Println("Servidor rodando na porta 5000")
	mux := http.NewServeMux()

	cssHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", cssHandler))

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)
	mux.HandleFunc("/note/create/store", noteStore)

	http.ListenAndServe(":5000", mux)
}
