package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DenisBuarque/goquicknotes.git/internal/handlers"
	"github.com/DenisBuarque/goquicknotes.git/internal/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	// conection data base
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:123456@localhost:5432/quicknotes")
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

	// Access repository Notes with conection in DB
	noteRepo := repositories.NewNoteRepository(dbpool)

	/*
		// Listing data notes
		notes, err := noteRepo.List()
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println(notes)
	*/

	/*
		// Get note id
		note, err := noteRepo.GetById(1)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Println(note)
	*/

	/*
		// Create Note
		note, err := noteRepo.Create("Javascript", "Aprendendo a framework NestJS", "#AAAACC")
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Print(note)
	*/

	// Update Note
	note, err := noteRepo.Update(1, "Laravel 10", "Aprendento o framework Php", "#FFAAAA")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(note)

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
