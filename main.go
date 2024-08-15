package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DenisBuarque/goquicknotes.git/internal/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	// conection data base
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	fmt.Println("Conexão com o banco realizada com sucesso.")

	// access file .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Erro loading .env file")
	}
	// access routes
	mux := http.NewServeMux()
	// Access file css
	cssHandler := http.FileServer(http.Dir("views/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", cssHandler))
	// Routes
	mux.HandleFunc("/", handlers.NewNoteHandler().NoteList)
	mux.HandleFunc("/note/view", handlers.NewNoteHandler().NoteView)
	mux.HandleFunc("/note/create", handlers.NewNoteHandler().NoteCreate)
	mux.HandleFunc("/note/create/store", handlers.NewNoteHandler().NoteStore)
	// Server
	port := os.Getenv("SERVER_PORT")
	fmt.Println("Servidor rodando na porta: " + port)
	http.ListenAndServe(port, mux)
}
