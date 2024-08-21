package repositories

import (
	"context"
	"time"

	"github.com/DenisBuarque/goquicknotes.git/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
	GetById(id int) (*models.Note, error)
	Create(title, content, color string) (*models.Note, error)
	Update(ID int, title, content, color string) (*models.Note, error)
}

func NewNoteRepository(dbpool *pgxpool.Pool) NoteRepository {
	return &connectDB{db: dbpool}
}

type connectDB struct {
	db *pgxpool.Pool
}

// implement from struct with interface
func (conn *connectDB) List() ([]models.Note, error) {
	var list []models.Note

	rows, err := conn.db.Query(context.Background(), "SELECT * FROM notes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, note)
	}

	return list, nil
}

func (conn *connectDB) GetById(id int) (*models.Note, error) {
	var note models.Note
	row := conn.db.QueryRow(context.Background(), `SELECT * FROM notes WHERE id = $1`, id)
	if err := row.Scan(&note.ID, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt); err != nil {
		return nil, err
	}
	return &note, nil
}

func (conn *connectDB) Create(title, content, color string) (*models.Note, error) {
	var note models.Note
	note.Title = title
	note.Content = content
	note.Color = color
	query := `INSERT INTO notes (title, content, color) VALUES ($1, $2, $3) RETURNING id, created_at`
	row := conn.db.QueryRow(context.Background(), query, note.Title, note.Content, note.Color)
	if err := row.Scan(&note.ID, &note.CreatedAt); err != nil {
		return nil, err
	}
	return &note, nil
}

func (conn *connectDB) Update(id int, title, content, color string) (*models.Note, error) {

	var note models.Note

	note.ID = id

	if len(title) > 0 {
		note.Title = title
	}

	if len(content) > 0 {
		note.Content = content
	}

	if len(color) > 0 {
		note.Color = color
	}

	note.UpdatedAt = time.Now()

	query := `UPDATE notes SET title=$1, content=$2, color=$3, updated_at=$4 WHERE id=$5`
	_, err := conn.db.Exec(context.Background(), query, note.Title, note.Content, note.Color, note.UpdatedAt, note.ID)
	if err != nil {
		return nil, err
	}
	return &note, nil

}
