package author

import (
	"fmt"
	"log"
	"projects/internal/model"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) AddAuthor(a *model.Author) (*model.Author, error) {
	// insert into authors (name, biography, address) values ('Leo Tolstoy', 'Russian', 'Russia')
	result := r.db.Create(&a)
	if result.Error != nil {
		log.Printf("AddAuthor: Failed to add author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}

	return a, nil
}

func (r *AuthorRepository) GetAuthors() ([]model.Author, error) {
	var authors []model.Author

	// select * from authors
	result := r.db.Find(&authors)
	if result.Error != nil {
		log.Printf("GetAuthors: Failed to get authors: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get authors: %v\n", result.Error)
	}

	return authors, nil
}

func (r *AuthorRepository) GetAuthorByID(authorID int) (*model.Author, error) {
	var author model.Author

	// select * from authors where author_id = authorID
	result := r.db.First(&author, authorID)
	if result.Error != nil {
		log.Printf("GetAuthorByID: Failed to get author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get author: %v\n", result.Error)
	}

	return &author, nil
}

func (r *AuthorRepository) UpdateAuthor(a *model.Author) (*model.Author, error) {
	// update authors set name = 'Leo Tolstoy', biography = 'Russian', address = 'Russia' where author_id = 1
	result := r.db.Model(&a).Updates(&a)
	if result.Error != nil {
		log.Printf("UpdateAuthor: Failed to update author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update author: %v\n", result.Error)
	}

	return a, nil
}

func (r *AuthorRepository) DeleteAuthor(authorID int) (int, error) {
	// delete from authors where author_id = authorID
	result := r.db.Delete(&model.Author{}, authorID)
	if result.Error != nil {
		log.Printf("DeleteAuthor: Failed to delete author: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete author: %v\n", result.Error)
	}

	return authorID, nil
}
