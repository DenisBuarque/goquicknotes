package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

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
	notes, err := nh.repository.List(r.Context())
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

	// verifica se a requisição ao banco demora mais de 1 min então cancela
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute)
	defer cancel()

	note, err := nh.repository.GetById(ctx, id)
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
		// informa o tipo de metodo obrigatório da requisição
		w.Header().Set("Allow", http.MethodPost)
		// escreve o tipo de erro no cabeçalho
		w.WriteHeader(405)
		fmt.Fprint(w, "Método não permitido")
		//http.Error(w, "Métpdo não permidito", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	color := r.PostForm.Get("color")

	note, err := nh.repository.Create(r.Context(), title, content, color)
	if err != nil {
		http.Error(w, "Erro ao criar nota.", http.StatusInternalServerError)
	}

	//http.Redirect(w, r, "/", http.StatusSeeOther)
	http.Redirect(w, r, fmt.Sprintf("/note/view?id=%d", note.ID), http.StatusSeeOther)
}
