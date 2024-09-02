package borrowDataBase

import (
	"fmt"
	"projects/internal/model"
)

type Service struct {
	b *BorrowRep
}

func NewService(b *BorrowRep) *Service {
	return &Service{b: b}
}
func (s *Service) CreateBorrow(b *model.Borrow) (*model.Borrow, error) {
	borrows, err := s.b.GetBorrows()
	if err != nil {
		return nil, err
	}
	if len(borrows) > 0 {
		for _, borrow := range borrows {
			if borrow.UserID == b.UserID && borrow.BookID == b.BorrowID {
				return nil, fmt.Errorf("such userID already borrowed this book")
			}
		}
	}
	return s.b.AddBorrow(b)
}
func (s *Service) ListBorrows() ([]model.Borrow, error) {
	borrows, err := s.b.GetBorrows()
	if err != nil {
		return nil, err
	} else if len(borrows) == 0 {
		return nil, fmt.Errorf("no borrows found")
	}
	return borrows, nil
}
func (s *Service) ListBorrowById(id int) (*model.Borrow, error) {
	b, err := s.b.GetBorrowById(id)
	if err != nil {
		return nil, err
	} else if b.BorrowID == 0 {
		return nil, fmt.Errorf("such borrowId doesn't exist")
	}
	return b, nil
}
func (s *Service) ListBorrowByUserId(UserID int) ([]model.Borrow, error) {
	borrows, err := s.b.GetBorrowsByUser(UserID)
	if err != nil {
		return nil, err
	} else if len(borrows) == 0 {
		return nil, fmt.Errorf("no borrow with such userID")
	}
	return borrows, nil
}
func (s *Service) EditBorrow(b *model.Borrow) (*model.Borrow, error) {
	bor, err := s.b.GetBorrowById(b.BorrowID)
	if err != nil {
		return nil, err
	} else if bor.BorrowID == 0 {
		return nil, fmt.Errorf("no borrow with such id")
	}
	return s.b.UpdateBorrow(b)
}
func (s *Service) RemoveBorrow(id int) (int, error) {
	bor, err := s.b.GetBorrowById(id)
	if err != nil {
		return 0, err
	} else if bor.BorrowID == 0 {
		return 0, fmt.Errorf("no borrow with such id")
	}
	return s.b.DeleteBorrowById(id)
}
