package repository

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) AddAuthor(author *model.Author) (*model.Author, error) {
	result := r.db.Create(&author)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}
	r.db.First(&author, author.AuthorID)
	return author, nil
}
