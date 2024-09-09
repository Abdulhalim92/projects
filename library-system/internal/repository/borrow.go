package repository

import (
	"fmt"
	"projects/internal/model"
)

func (r *Repository) AddBorrow(bor *model.Borrow) (*model.Borrow, error) {
	err := r.db.Create(bor).Error
	if err != nil {
		return nil, fmt.Errorf("error while adding a borrow: %v", err)
	}
	return bor, nil
}
func (r *Repository) GetBorrows() ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := r.db.Find(&borrows).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting borrows: %v", err)
	}
	return borrows, nil
}
func (r *Repository) GetBorrowById(id int) (*model.Borrow, error) {
	var borrow model.Borrow
	err := r.db.Table("borrow").Where("borrow_id = ?", id).Select("borrow_id", "user_id", "book_id", "borrowdate", "returndate").Scan(&borrow).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a borrow")
	}
	return &borrow, nil
}
func (r *Repository) GetBorrowsByUser(UserID int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := r.db.Table("borrow").Where("user_id = ?", UserID).Scan(&borrows).Error
	if err != nil {
		return nil, fmt.Errorf("error getting borrows by userID: %v", err)
	}
	return borrows, nil
}
func (r *Repository) UpdateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	err := r.db.Table("borrow").Updates(borrow).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a borrow: %v", err)
	}
	return borrow, nil
}
func (r *Repository) DeleteBorrowById(BorrowID int) (int, error) {
	err := r.db.Delete(&model.Borrow{}, BorrowID).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting a borrow")
	}
	return BorrowID, nil
}
