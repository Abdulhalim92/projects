package AuthorDataBase

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type AuthorRep struct {
	db *gorm.DB
}

func NewAuthorRep(db *gorm.DB) *AuthorRep {
	return &AuthorRep{db}
}
func (a *AuthorRep) AddAuthor(author *model.Author) (*model.Author, error) {
	err := a.db.Table("authors").Create(author).Error
	if err != nil {
		return nil, fmt.Errorf("error adding a book: %v", err)
	}
	return author, nil
}
func (a *AuthorRep) GetAuthors() ([]model.Author, error) {
	var authors []model.Author
	err := a.db.Find(&authors).Error
	if err != nil {
		return nil, fmt.Errorf("error getting authors: %v", err)
	}
	return authors, nil
}
func (a *AuthorRep) GetAuthorById(id int) (*model.Author, error) {
	var author model.Author
	err := a.db.Table("authors").Where("author_id = ?", id).Select("author_id", "name", "biography").Scan(&author).Error
	if err != nil {
		return nil, fmt.Errorf("error getting an author: %v", err)
	}
	return &author, nil
}
func (a *AuthorRep) UpdateAuthor(author *model.Author) (*model.Author, error) {
	err := a.db.Table("authors").Updates(author).Error
	if err != nil {
		return nil, fmt.Errorf("error updating an author: %v", err)
	}
	return author, nil
}
func (a *AuthorRep) DeleteAuthor(id int) (int, error) {
	err := a.db.Table("authors").Delete(&model.Author{}, id).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting an author: %v", err)
	}
	return id, nil
}
