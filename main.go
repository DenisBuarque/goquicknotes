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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env file")
	}

	port := os.Getenv("SERVER_PORT")

	fmt.Println("Servidor rodando na porta: " + port)
	mux := http.NewServeMux()

	cssHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", cssHandler))

	mux.HandleFunc("/", handlers.NoteList)
	mux.HandleFunc("/note/view", handlers.NoteView)
	mux.HandleFunc("/note/create", handlers.NoteCreate)
	mux.HandleFunc("/note/create/store", handlers.NoteStore)

	http.ListenAndServe(port, mux)
}
