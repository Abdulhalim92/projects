package internal

import (
	"fmt"
	"log"
	"projects/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// BookRepository

func (r *Repository) AddBook(b *model.Book) (*model.Book, error) {
	result := r.db.Create(&b)
	// insert into books (title, author_id) values ('War and Peace', 1)
	if result.Error != nil {
		log.Printf("Addbook: Failed to add book: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add book: %v\n", result.Error)
	}

	return b, nil
}

func (r *Repository) GetBooks() ([]model.Book, error) {
	var b []model.Book

	// select * from books;
	err := r.db.Find(&b).Error
	if err != nil {
		log.Printf("GetBooks: Failed to get books: %v\n", err)
		return nil, fmt.Errorf("Cannon find books with error: %v", err)
	}

	return b, nil
}

func (r *Repository) GetBookByID(bookID int) (*model.Book, error) {
	var b model.Book

	// select * from books where book_id = bookID
	err := r.db.Where("book_id = ?", bookID).Find(&b).Error
	if err != nil {
		log.Printf("GetBookByID: Failed to get book: %v\n", err)
		return nil, fmt.Errorf("Cannot find book with error: %v", err)
	}

	return &b, nil
}

func (r *Repository) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	var b []model.Book

	// select * from books where author_id = authorID
	err := r.db.Where("author_id = ?", authorID).Find(&b).Error
	if err != nil {
		log.Printf("GetBooksByAuthor: Failed to get books: %v\n", err)
		return nil, fmt.Errorf("Cannot find books by authorID with error: %v", err)
	}

	return b, nil
}

func (r *Repository) UpdateBook(b *model.Book) (*model.Book, error) {
	// update books set title = 'War and Peace', author_id = 1 where book_id = 1
	result := r.db.Model(&b).Updates(&b)
	if result.Error != nil {
		log.Printf("UpdateBook: Failed to update book: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update book: %v\n", result.Error)
	}

	return b, nil
}

func (r *Repository) DeleteBook(bookID int) (int, error) {
	// delete from books where book_id = bookID returning book_id
	result := r.db.Delete(&model.Book{}, bookID)
	if result.Error != nil {
		log.Printf("DeleteBook: Failed to delete book: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete book: %v\n", result.Error)
	}

	return bookID, nil
}

// UserRepository

func (r *Repository) AddUser(u *model.User) (*model.User, error) {
	// insert into users (username, password) values ('admin', 'admin')
	result := r.db.Create(&u)
	if result.Error != nil {
		log.Printf("AddUser: Failed to add user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add user: %v\n", result.Error)
	}

	return u, nil
}

func (r *Repository) GetUsers() ([]model.User, error) {
	var users []model.User

	// select * from users
	result := r.db.Find(&users)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get users: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get users: %v\n", result.Error)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id int) (*model.User, error) {
	var user model.User

	// select * from users where user_id = id
	result := r.db.First(&user, id)
	if result.Error != nil {
		log.Printf("GetUserByID: Failed to get user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get user: %v\n", result.Error)
	}
	return &user, nil
}

func (r *Repository) UpdateUser(u *model.User) (*model.User, error) {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&u).Updates(&u)
	if result.Error != nil {
		log.Printf("UpdateUser: Failed to update user: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update user: %v\n", result.Error)
	}

	return u, nil
}

func (r *Repository) DeleteUser(id int) (int, error) {
	// delete from users where user_id = id
	result := r.db.Delete(&model.User{}, id)
	if result.Error != nil {
		log.Printf("DeleteUser: Failed to delete user: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete user: %v\n", result.Error)
	}

	return id, nil
}

// AuthorRepository

func (r *Repository) AddAuthor(a *model.Author) (*model.Author, error) {
	// insert into authors (name, biography, address) values ('Leo Tolstoy', 'Russian', 'Russia')
	result := r.db.Create(&a)
	if result.Error != nil {
		log.Printf("AddAuthor: Failed to add author: %v\n", result.Error)
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

func (r *Repository) UpdateAuthor(a *model.Author) (*model.Author, error) {
	// update authors set name = 'Leo Tolstoy', biography = 'Russian', address = 'Russia' where author_id = 1
	result := r.db.Model(&a).Updates(&a)
	if result.Error != nil {
		log.Printf("UpdateAuthor: Failed to update author: %v\n", result.Error)
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
