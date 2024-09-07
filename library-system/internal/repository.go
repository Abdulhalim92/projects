package internal

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"projects/internal/model"
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
