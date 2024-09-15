package service

import (
	"fmt"
	"gorm.io/gorm"
	"projects/internal/model"
)

func (s *Service) ListBorrows() ([]*model.Borrow, error) {
	return s.Repository.GetBorrows()
}

func (s *Service) ListFilteredBorrows(filter *model.BorrowFilter) ([]*model.Borrow, error) {
	return s.Repository.GetBorrowByFilter(filter)
}

func (s *Service) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	book, err := s.Repository.GetBookByID(borrow.BookID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if book == nil {
		return nil, fmt.Errorf("Book with ID %d does not exist\n", borrow.BookID)
	}
	user, err := s.Repository.GetUserByID(borrow.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if user == nil {
		return nil, fmt.Errorf("Book with ID %d does not exist\n", borrow.BookID)
	}
	filter := model.BorrowFilter{UserID: borrow.UserID, BookID: borrow.BookID, WasReturned: false, CountOnPage: 1, Page: 1}
	borrows, err := s.Repository.GetBorrowByFilter(&filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to get borrows by filter\n")
	}
	fmt.Println(filter)
	if len(borrows) > 0 {
		return nil, fmt.Errorf("The book with id %d was already borrowed by user with id %d\n", borrow.BookID, borrow.UserID)
	}
	return s.Repository.AddBorrow(borrow)
}

func (s *Service) ReturnBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	book, err := s.Repository.GetBookByID(borrow.BookID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if book == nil {
		return nil, fmt.Errorf("Book with ID %d does not exist\n", borrow.BookID)
	}
	user, err := s.Repository.GetUserByID(borrow.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if user == nil {
		return nil, fmt.Errorf("Book with ID %d does not exist\n", borrow.BookID)
	}
	return s.Repository.ReturnBorrow(borrow)
}
