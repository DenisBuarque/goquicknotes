package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/home.html")
	if err != nil {
		http.Error(w, "Error ao carregar página.", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func noteView(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id não encontrado", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("views/templates/show.html")
	if err != nil {
		http.Error(w, "Erro ao carregar página show", http.StatusInternalServerError)
		return
	}
	t.Execute(w, id)
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("/", noteList)
	mux.HandleFunc("/note/view", noteView)
	mux.HandleFunc("/note/create", noteCreate)

	http.ListenAndServe(":5000", mux)
}
