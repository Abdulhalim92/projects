package author

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepo{db: db}
}

func (r *AuthorRepo) AddAuthor(a *model.Author) (*model.Author, error) {
	// insert into users (username, password) values ('admin', 'admin')
	result := r.db.Create(&a)
	if result.Error != nil {
		log.Printf("AddAuthor: Failed to add author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add author: %v\n", result.Error)
	}

	return a, nil
}

func (r *AuthorRepo) GetAuthors() ([]model.Author, error) {
	var authors []model.Author

	// select * from users
	result := r.db.Find(&authors)
	if result.Error != nil {
		log.Printf("GetUsers: Failed to get authors: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get authors: %v\n", result.Error)
	}
	return authors, nil
}

func (r *AuthorRepo) GetAuthorByID(id int) (*model.Author, error) {
	var author model.Author

	// select * from users where user_id = id
	result := r.db.First(&author, id)
	if result.Error != nil {
		log.Printf("GetAuthorByID: Failed to get author: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get author: %v\n", result.Error)
	}
	return &author, nil
}

func (r *AuthorRepo) UpdateAuthor(a *model.Author) error {
	// update users set username = 'admin', password = 'admin' where user_id = 1
	result := r.db.Model(&a).Updates(&a)
	if result.Error != nil {
		log.Printf("UpdateAuthor: Failed to update author: %v\n", result.Error)
		return fmt.Errorf("Failed to update author: %v\n", result.Error)
	}

	return nil
}

func (r *AuthorRepo) DeleteAuthor(id int) error {
	// delete from users where user_id = id
	result := r.db.Delete(&model.Author{}, id)
	if result.Error != nil {
		log.Printf("DeleteAuthor: Failed to delete author: %v\n", result.Error)
		return fmt.Errorf("Failed to delete author: %v\n", result.Error)
	}

	return nil
}
