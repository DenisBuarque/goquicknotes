package models

import (
	"time"
)

type Note struct {
	ID        string    `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Color     string    `json:"color" bson:"color"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"created_at"`
}

/*
type Note struct {
	ID        pgtype.Numeric
	Title     pgtype.Text
	Content   pgtype.Text
	Color     pgtype.Text
	CreatedAt pgtype.Date
	UpdatedAt pgtype.Date
	//UpdatedAt   sql.NullString
}
*/
