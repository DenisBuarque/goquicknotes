package repositories

import (
	"context"

	"github.com/DenisBuarque/goquicknotes.git/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
	GetById(id int) (*models.Note, error)
	Create(title, content, color string) (*models.Note, error)
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
