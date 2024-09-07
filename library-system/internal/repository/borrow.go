package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) AddBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	// insert into borrows (user_id, book_id, borrow_date, return_date) values (user_id, book_id, borrow_date, return_date) returning borrow_id
	result := r.db.Create(&borrow)
	if result.Error != nil {
		log.Printf("AddBorrow: Failed to add borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to add borrow: %v\n", result.Error)
	}

	return borrow, nil
}

func (r *Repository) GetBorrows() ([]model.Borrow, error) {
	// select * from borrows
	var borrows []model.Borrow
	result := r.db.Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrows: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

func (r *Repository) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	var borrows []model.Borrow

	// select * from borrows where user_id = userID
	result := r.db.Where("user_id = ?", userID).Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrowsByUser: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

func (r *Repository) GetBorrowsByBook(bookID int) ([]model.Borrow, error) {
	var borrows []model.Borrow

	// select * from borrows where book_id = bookID
	result := r.db.Where("book_id = ?", bookID).Find(&borrows)
	if result.Error != nil {
		log.Printf("GetBorrowsByBook: Failed to get borrows: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrows: %v\n", result.Error)
	}

	return borrows, nil
}

func (r *Repository) GetBorrowByID(borrowID int) (*model.Borrow, error) {
	var borrow model.Borrow

	// select * from borrows where borrow_id = borrowID
	result := r.db.Where("borrow_id = ?", borrowID).Find(&borrow)
	if result.Error != nil {
		log.Printf("GetBorrowByID: Failed to get borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrow: %v\n", result.Error)
	}

	return &borrow, nil
}

func (r *Repository) UpdateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	// update borrows set return_date = now() where borrow_id = borrowID
	result := r.db.Model(&borrow).Updates(&borrow)
	if result.Error != nil {
		log.Printf("UpdateBorrow: Failed to update borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update borrow: %v\n", result.Error)
	}

	return borrow, nil
}

func (r *Repository) DeleteBorrow(borrowID int) (int, error) {
	// delete from borrows where borrow_id = borrowID returning borrow_id
	result := r.db.Where("borrow_id = ?", borrowID).Delete(&model.Borrow{})
	if result.Error != nil {
		log.Printf("DeleteBorrow: Failed to delete borrow: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete borrow: %v\n", result.Error)
	}

	return borrowID, nil
}

func (r *Repository) GetBorrowsByUserAndBook(userID, bookID int) (*model.Borrow, error) {

	var borrow model.Borrow
	// select * from borrows where user_id = userID and book_id = bookID
	result := r.db.Where("user_id = ? AND book_id = ?", userID, bookID).Find(&borrow)
	if result.Error != nil {
		log.Printf("GetBorrowsByUserAndBook: Failed to get borrow: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to get borrow: %v\n", result.Error)
	}

	return &borrow, nil
}
