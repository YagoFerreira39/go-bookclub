package entities

import "github.com/google/uuid"

type BookEntity struct {
	ID        string
	Name      string
	Author    string
	ISBN      string
	Published float64
}

func NewBookEntity(id *string, name, author, isbn string, published float64) *BookEntity {
	newID := uuid.New().String()
	if id == nil {
		id = &newID
	}

	entity := &BookEntity{
		ID: *id,
		Name: name,
		Author: author,
		ISBN: isbn,
		Published: published,
	}

	return entity
}
