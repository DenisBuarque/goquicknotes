package repositories

import (
	"context"

	"github.com/DenisBuarque/goquicknotes.git/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NoteRepository interface {
	List() ([]models.Note, error)
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
