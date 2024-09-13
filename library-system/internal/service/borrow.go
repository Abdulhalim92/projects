package service

import (
	"errors"
	"fmt"
	"projects/internal/model"
	"time"
)

var (
	ErrBookNotAvailable = errors.New("book not available")
	ErrBorrowsNotFound  = errors.New("no borrows found")
)

func (s *Service) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	bookAvailable, err := s.Repository.IsBookAvailable(borrow.BookID)
	if err != nil {
		return nil, err
	}
	// Проверяем, доступна ли книга для выдачи
	if !bookAvailable {
		return nil, ErrBookNotAvailable
	}
	// Установка даты выдачи и возврата
	now := time.Now()
	borrow.BorrowDate = &now

	return s.Repository.AddBorrow(borrow)
}

func (s *Service) ReturnBook(borrowID int) error {
	// Получаем информацию о выдаче книги по ID
	borrowByID, err := s.Repository.GetBorrowByID(borrowID)
	if err != nil {
		return err
	}
	if borrowByID == nil {
		return fmt.Errorf("borrow not found with ID %d", borrowID)
	}
	if borrowByID.ReturnDate != nil {
		return fmt.Errorf("book already returned with ID %d", borrowID)
	}
	// Получаем информацию о книге по ID
	bookByBorrow, err := s.Repository.GetBookByBorrow(borrowID)
	if err != nil {
		return err
	}
	if bookByBorrow == nil {
		return fmt.Errorf("book not found with borrow ID %d", borrowID)
	}

	return s.Repository.ReturnBook(borrowID)
}

func (s *Service) GetBorrows() ([]model.Borrow, error) {
	borrows, err := s.Repository.GetBorrows()
	if err != nil {
		return nil, err
	}
	if len(borrows) == 0 {
		return nil, ErrBorrowsNotFound
	}

	return borrows, nil
}

func (s *Service) GetBorrowByID(borrowID int) (*model.Borrow, error) {
	borrowByID, err := s.Repository.GetBorrowByID(borrowID)
	if err != nil {
		return nil, err
	}
	if borrowByID == nil {
		return nil, fmt.Errorf("borrow not found with ID %d", borrowID)
	}

	return borrowByID, nil
}

func (s *Service) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	borrowsByUser, err := s.Repository.GetBorrowsByUser(userID)
	if err != nil {
		return nil, err
	}
	if len(borrowsByUser) == 0 {
		return nil, fmt.Errorf("no borrows found with user ID %d", userID)
	}

	return borrowsByUser, nil
}

func (s *Service) GetBorrowsByBook(bookID int) ([]model.Borrow, error) {
	borrowsByBook, err := s.Repository.GetBorrowsByBook(bookID)
	if err != nil {
		return nil, err
	}
	if len(borrowsByBook) == 0 {
		return nil, fmt.Errorf("no borrows found with book ID %d", bookID)
	}

	return borrowsByBook, nil
}

func (s *Service) GetBorrowsByUserAndBook(userID, bookID int) (*model.Borrow, error) {
	borrowByUserAndBook, err := s.Repository.GetBorrowsByUserAndBook(userID, bookID)
	if err != nil {
		return nil, err
	}
	if borrowByUserAndBook == nil {
		return nil, fmt.Errorf("borrow not found with user ID %d and book ID %d", userID, bookID)
	}

	return borrowByUserAndBook, nil
}
