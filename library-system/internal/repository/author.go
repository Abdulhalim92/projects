package repository

import (
	"fmt"
	"projects/internal/model"
)

// Author Handling Methods

func (r *Repository) GetAuthors() ([]*model.Author, error) {
	var authors []*model.Author
	result := r.db.Find(&authors)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get authors: %v\n", result.Error)
	}
	return authors, nil
}

func (r *Repository) GetAuthorByID(id int) (*model.Author, error) {
	var author model.Author
	result := r.db.First(&author, id)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to get author by ID: %v\n", result.Error)
	}
	return &author, nil
}

func (r *Repository) AddAuthor(author *model.Author) (*model.Author, error) {
	result := r.db.Create(&author)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}
	result = r.db.First(&author, author.AuthorID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}
	return author, nil
}

func (r *Repository) EditAuthor(author *model.Author) (*model.Author, error) {

	result := r.db.Save(&author)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to edit author: %v\n", result.Error)
	}
	var a *model.Author
	result = r.db.First(&a, author.AuthorID)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to edit author: %v\n", result.Error)
	}
	return author, nil

}
