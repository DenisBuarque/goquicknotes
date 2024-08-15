package connnection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Post struct {
	id      int
	title   string
	content string
	author  string
}

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

	//createTable()
	//insertPost()
	//postId()
	postsAll()
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
	title := "Post 3"
	content := "Conteúdo do post 3"
	author := "Denis"

	query := `INSERT INTO posts (title, content, author) VALUES ($1, $2, $3)`
	_, err := conn.Exec(context.Background(), query, title, content, author)
	if err != nil {
		panic(err)
	}
	fmt.Println("Post inserido com sucesso.")
}

func postId() {
	id := 1
	var title, content, author string
	query := "SELECT title, content, author FROM posts WHERE id=$1"
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&title, &content, &author)
	if err == pgx.ErrNoRows {
		fmt.Println("Post not found for id: ", id)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("Titulo: %s, Conteúdo: %s, Autor: %s \n", title, content, author)
}

func postsAll() {
	query := "SELECT * FROM posts"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.id, &post.title, &post.content, &post.author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	for _, post := range posts {
		fmt.Println(post)
	}
}
