package borrowDataBase

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

type BorrowRep struct {
	db *gorm.DB
}

func NewBorrowRep(db *gorm.DB) *BorrowRep {
	return &BorrowRep{db}
}
func (b *BorrowRep) AddBorrow(bor *model.Borrow) (*model.Borrow, error) {
	err := b.db.Create(bor).Error
	if err != nil {
		return nil, fmt.Errorf("error while adding a borrow: %v", err)
	}
	return bor, nil
}
func (b *BorrowRep) GetBorrows() ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := b.db.Find(&borrows).Error
	if err != nil {
		return nil, fmt.Errorf("error while getting borrows: %v", err)
	}
	return borrows, nil
}
func (b *BorrowRep) GetBorrowById(id int) (*model.Borrow, error) {
	var borrow model.Borrow
	err := b.db.Table("borrow").Where("borrow_id = ?", id).Select("borrow_id", "user_id", "book_id", "borrowdate", "returndate").Scan(&borrow).Error
	if err != nil {
		return nil, fmt.Errorf("error getting a borrow")
	}
	return &borrow, nil
}
func (b *BorrowRep) GetBorrowsByUser(UserID int) ([]model.Borrow, error) {
	var borrows []model.Borrow
	err := b.db.Table("borrow").Where("user_id = ?", UserID).Scan(&borrows).Error
	if err != nil {
		return nil, fmt.Errorf("error getting borrows by userID: %v", err)
	}
	return borrows, nil
}
func (b *BorrowRep) UpdateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	err := b.db.Table("borrow").Updates(borrow).Error
	if err != nil {
		return nil, fmt.Errorf("error while updating a borrow: %v", err)
	}
	return borrow, nil
}
func (b *BorrowRep) DeleteBorrowById(BorrowID int) (int, error) {
	err := b.db.Delete(&model.Borrow{}, BorrowID).Error
	if err != nil {
		return 0, fmt.Errorf("error deleting a borrow")
	}
	return BorrowID, nil
}
