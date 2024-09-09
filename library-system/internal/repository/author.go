package repository

import (
	"fmt"
	"projects/internal/model"
)

func (r *Repository) AddAuthor(author *model.Author) (*model.Author, error) {
	err := r.db.Table("authors").Create(author).Error
	if err != nil {
		return nil, fmt.Errorf("error adding a book: %v", err)
	}
	return author, nil
}
func (r *Repository) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := r.db.Find(&authors).Error
	if err != nil {
		return nil, fmt.Errorf("error getting authors: %v", err)
	}
	return authors, nil
}
func (r *Repository) GetAuthorById(id int) (*model.Author, error) {
	var author model.Author
	err := r.db.Table("authors").Where("author_id = ?", id).Select("author_id", "name", "biography").Scan(&author).Error
	if err != nil {
		return nil, fmt.Errorf("error getting an author: %v", err)
	}
	return &author, nil
}
func (r *Repository) UpdateAuthor(author *model.Author) (*model.Author, error) {
	err := r.db.Table("authors").Updates(author).Error
	if err != nil {
		return nil, fmt.Errorf("error updating an author: %v", err)
	}
	return author, nil
}
func (r *Repository) DeleteAuthor(id int) (int, error) {
	err := r.db.Table("authors").Delete(&model.Author{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting an author: %v", err)
	}
	return id, nil
}
