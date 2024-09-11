package repository

import "gorm.io/gorm"

//import "time"

type BorrowsRepo struct {
	db *gorm.DB
}

func NewBorrowsRepo(db *gorm.DB) *BorrowsRepo {
	return &BorrowsRepo{db: db}
}
