package borrow

import (
	"fmt"
	"library-system/internal/model"
	"time"
)

type Service struct {
	br BorrowRepo
}

func NewService(b BorrowRepo) *Service {
	return &Service{b}
}

func (s *Service) CreateBorrow(borrowID, userID, bookID int, borrowDate, returnDate time.Time) (*model.Borrow, error) {
	borrow := model.Borrow{BorrowId: borrowID, UserId: userID, BookId: bookID, BorrowDate: borrowDate, ReturnDate: returnDate}
	return s.br.AddBorrow(&borrow)
}

func (s *Service) ListBorrows() ([]model.Borrow, error) {
	borrows, err := s.br.GetBorrows()
	if err != nil {
		return nil, fmt.Errorf("Error when listing the borrows: %e", err)
	}

	return borrows, nil
}

func (s *Service) FindBorrow(id int) (*model.Borrow, error) {
	borrow, err := s.br.GetBorrowByID(id)
	if err != nil {
		return nil, fmt.Errorf("Error occured when retrieiving borrow with id:%d\n%e", id, err)
	}

	return borrow, nil
}

func (s *Service) FindBorrowsByUsername(username string) ([]model.Borrow, error) {
	user, err := s.FindBorrowByUsername(username)
	if err != nil {
		return nil, err
	}

	borrows, err := s.br.GetBorrowsByUser(user.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't get borrows of user with id:%d\n%e", user.UserId, err)
	}

	return borrows, nil
}

func (s *Service) EditBorrow(id, userID, bookID int, borrowDate, returnDate time.Time) error {
	borrow := model.Borrow{
		BorrowId:   id,
		UserId:     userID,
		BookId:     bookID,
		BorrowDate: borrowDate,
		ReturnDate: returnDate,
	}
	err := s.br.UpdateBorrow(&borrow)
	if err != nil {
		return fmt.Errorf("Error occured when editing borrow with id:%d\n%e", id, err)
	}
	return nil
}

func (s *Service) RemoveBorrow(id int) bool {
	err := s.br.DeleteBorrow(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (s *Service) FindBorrowByUsername(name string) (*model.User, error) {
	var user model.User
	err := s.br.db.First(&user, "username = ?", name).Error
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Couldn't find borrow with username:%s\n%e", name, err)
	}

	return &user, nil
}
