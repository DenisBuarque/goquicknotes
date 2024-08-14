package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DenisBuarque/goquicknotes.git/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// access file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env file")
	}

	// message acess on server
	port := os.Getenv("SERVER_PORT")
	fmt.Println("Servidor rodando na porta: " + port)
	// access routes
	mux := http.NewServeMux()
	// access file css
	cssHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", cssHandler))
	// routes navigation
	mux.HandleFunc("/", handlers.NewNoteHandler().NoteList)
	mux.HandleFunc("/note/view", handlers.NewNoteHandler().NoteView)
	mux.HandleFunc("/note/create", handlers.NewNoteHandler().NoteCreate)
	mux.HandleFunc("/note/create/store", handlers.NewNoteHandler().NoteStore)
	// create server
	http.ListenAndServe(port, mux)
}
