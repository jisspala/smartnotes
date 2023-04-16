package models

import (
	"time"
)

type Note struct {
	ID        int    `json:"id" gorm:"primarykey"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MultiID struct {
	IDs []int `json:"ids"`
}
