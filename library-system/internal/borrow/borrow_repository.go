package borrow

import (
	"fmt"
	"library-system/internal/model"
	"log"

	"gorm.io/gorm"
)

type BorrowRepo struct {
	db *gorm.DB
}

func NewBorrowRepo(db *gorm.DB) *BorrowRepo {
	return &BorrowRepo{db: db}
}

func (r *BorrowRepo) AddBorrow(b *model.Borrow) (*model.Borrow, error) {
	result := r.db.Create(&b)
	// insert into borrows (title, author_id) values ('War and Peace', 1)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add borrow: %v\n", result.Error)
	}

	return b, nil
}

func (r *BorrowRepo) GetBorrows() ([]model.Borrow, error) {
	var borrows []model.Borrow

	err := r.db.Find(&borrows).Error
	if err != nil {
		return nil, err
	}
	return borrows, nil
}

func (r *BorrowRepo) GetBorrowByID(borrowID int) (*model.Borrow, error) {
	var borrow model.Borrow
	err := r.db.First(&borrow, "id = ?", borrowID).Error
	if err != nil {
		return nil, err
	}

	return &borrow, nil
}

func (r *BorrowRepo) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	var borrow []model.Borrow
	err := r.db.Find(&borrow, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}

	return borrow, nil
}

func (r *BorrowRepo) UpdateBorrow(b *model.Borrow) error {
	result := r.db.Model(&b).Updates(&b)
	if result.Error != nil {
		log.Printf("UpdateBorrow: Failed to update borrow: %v\n", result.Error)
		return fmt.Errorf("Failed to update borrow: %v\n", result.Error)
	}

	return nil
}

func (r *BorrowRepo) DeleteBorrow(borrowID int) error {
	borrow, err := r.GetBorrowByID(borrowID)
	if err != nil {
		return err
	}
	err = r.db.Delete(&borrow).Error
	if err != nil {
		return err
	}

	return nil
}
