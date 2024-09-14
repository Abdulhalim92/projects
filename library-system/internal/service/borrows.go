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
	return s.Repository.AddBorrow(borrow)

}
