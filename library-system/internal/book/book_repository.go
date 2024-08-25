package book

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type Repo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) AddBook(b *model.Book) (*model.Book, error) {
	result := r.db.Create(&b)
	// insert into books (title, author_id) values ('War and Peace', 1)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}

	return b, nil
}

func (r *Repo) GetBooks() ([]model.Book, error) {

}

func (r *Repo) GetBookByID(bookID int) (*model.Book, error) {

}

func (r *Repo) GetBooksByAuthor(authorID int) ([]model.Book, error) {

}

func (r *Repo) UpdateBook(b *model.Book) (*model.Book, error) {

}

func (r *Repo) DeleteBook(bookID int) (int, error) {

}
