package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/DenisBuarque/goquicknotes.git/internal/repositories"
)

type noteHandler struct {
	repository repositories.NoteRepository
}

func NewNoteHandler(repo repositories.NoteRepository) *noteHandler {
	return &noteHandler{repository: repo}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/layoutBase.html",
		"views/templates/pages/home.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Error ao carregar página.", http.StatusInternalServerError)
		return
	}

	// access note repository
	notes, err := nh.repository.List()
	if err != nil {
		http.Error(w, "Error ao listar dados.", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "layoutBase", listNoteResponsedto(notes))
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"views/templates/layoutBase.html",
		"views/templates/pages/show.html",
	}

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		http.Error(w, "id não encontrado", http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Erro ao carregar página show", http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		panic(err)
	}

	note, err := nh.repository.GetById(id)
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "layoutBase", noteResponseDto(note))
}

func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request) {
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

func (nh *noteHandler) NoteStore(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost) // informa o tipo de metodo obrogatório da requisição

		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido")
		//http.Error(w, "Métpdo não permidito", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Criar uma noto.")
}
