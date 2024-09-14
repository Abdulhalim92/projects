package repository

import (
	"fmt"
	"projects/internal/model"
)

//import "time"

func (r *Repository) GetBorrows() ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	result := r.db.Find(&borrows)
	if result.Error != nil {
		fmt.Errorf("Failed to get borrows")
		return nil, result.Error
	}
	return borrows, nil
}

func (r *Repository) GetBorrowByFilter(filter *model.BorrowFilter) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	var query = r.db

	if filter.BorrowID > 0 {
		query = query.Where("borrow_id = ?", filter.BorrowID)
	}
	if filter.UserID > 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.BookID > 0 {
		query = query.Where("book_id = ?", filter.BookID)
	}
	if filter.WasReturned {
		query = query.Where("return_date IS NOT NULL")
	} else {
		query = query.Where("return_date IS NULL")
	}
	if !filter.DateFrom.IsZero() {
		query = query.Where("borrow_date >= ?", filter.DateFrom)
	}
	if !filter.DateTo.IsZero() {
		query = query.Where("borrow_date <= ?", filter.DateTo)
	}
	query.Offset((filter.Page - 1) * filter.CountOnPage).Limit(filter.CountOnPage).Find(&borrows)

	if query.Error != nil {
		fmt.Errorf("Failed to get filtered borrows")
		return nil, query.Error
	}

	return borrows, nil
}

func (r *Repository) AddBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	result := r.db.Create(&borrow)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to add borrow: %v\n", result.Error)
	}
	return borrow, nil

}
