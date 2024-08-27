package handlers

import "github.com/DenisBuarque/goquicknotes.git/internal/models"

type NoteResponseDto struct {
	ID      int
	Title   string
	Content string
	Color   string
}

func noteResponseDto(note *models.Note) (res NoteResponseDto) {
	res.ID = note.ID
	res.Title = note.Title
	res.Content = note.Content
	res.Color = note.Color
	return
}

func listNoteResponsedto(notes []models.Note) (res []NoteResponseDto) {
	for _, note := range notes {
		res = append(res, noteResponseDto(&note))
	}
	return
}
