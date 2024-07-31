package main

import (
	"fmt"
	"net/http"
)

func noteList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Lista de notas.")
}

func noteView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Exibir uma nota.")
}

func noteCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {

		w.Header().Set("Allow", "POST") // informa o tipo de metodo da requisição

		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido")
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
