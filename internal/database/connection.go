package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {
	var err error

	dbURL := "postgres://postgres:123456@localhost:5432/quicknotes"
	conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso.")
	defer conn.Close(context.Background())

	createTable()
	insertPost()
}

func createTable() {
	query := `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title varchar(50) NOT NULL,
			content text NULL,
			author varchar(50) NOT NULL
		);`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tabela posts criada com sucesso.")
}

func insertPost() {
	title := "Post 1"
	content := "Conteúdo do post 1"
	author := "Denis"

	query := fmt.Sprintf(`INSERT INTO (title, content, author) VALUES ('%s', '%s', '%s')`, title, content, author)
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Post inserido com sucesso.")
}
