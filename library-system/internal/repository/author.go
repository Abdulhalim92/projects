package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) AddAuthor(a *model.Author) (*model.Author, error) {
	// insert into authors (name, biography, address) values ('Leo Tolstoy', 'Russian', 'Russia')
	result := r.db.Create(&a)
	if result.Error != nil {
		log.Printf("CreateAuthor: Failed to add author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}

	return a, nil
}

func (r *Repository) GetAuthors() ([]model.Author, error) {
	var authors []model.Author

	// select * from authors
	result := r.db.Find(&authors)
	if result.Error != nil {
		log.Printf("GetAuthors: Failed to get authors: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get authors: %v\n", result.Error)
	}

	return authors, nil
}

func (r *Repository) GetAuthorByID(authorID int) (*model.Author, error) {
	var author model.Author

	// select * from authors where author_id = authorID
	result := r.db.First(&author, authorID)
	if result.Error != nil {
		log.Printf("GetAuthorByID: Failed to get author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get author: %v\n", result.Error)
	}

	return &author, nil
}

func (r *Repository) GetAuthorByName(name string) (*model.Author, error) {
	var author model.Author

	// select * from authors where name = name
	result := r.db.Where("name = ?", name).First(&author)
	if result.Error != nil {
		log.Printf("GetAuthorByName: Failed to get author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get author: %v\n", result.Error)
	}

	return &author, nil
}

func (r *Repository) UpdateAuthor(a *model.Author) (*model.Author, error) {
	// update authors set name = 'Leo Tolstoy', biography = 'Russian', address = 'Russia' where author_id = 1
	result := r.db.Model(&a).Updates(&a)
	if result.Error != nil {
		log.Printf("EditAuthor: Failed to update author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update author: %v\n", result.Error)
	}

	return a, nil
}

func (r *Repository) DeleteAuthor(authorID int) (int, error) {
	// delete from authors where author_id = authorID
	result := r.db.Delete(&model.Author{}, authorID)
	if result.Error != nil {
		log.Printf("DeleteAuthor: Failed to delete author: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete author: %v\n", result.Error)
	}

	return authorID, nil
}
