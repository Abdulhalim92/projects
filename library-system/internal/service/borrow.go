package service

import (
	"errors"
	"fmt"
	"projects/internal/model"
	"time"
)

var (
	ErrBookNotAvailable    = errors.New("book not available")
	ErrBorrowsNotFound     = errors.New("no borrows found")
	ErrBorrowNotFound      = errors.New("borrow not found")
	ErrBookAlreadyReturned = errors.New("book already returned")
	ErrDifferentUser       = errors.New("different user")
	ErrUserNotFound        = errors.New("user not found")
)

func (s *Service) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(borrow.UserID)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			//return nil, fmt.Errorf("user with id %d doesn't exist", borrow.UserID)
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Проверяем, имеется ли книга в наличии
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

func (s *Service) ReturnBook(userID int, borrowID int) error {
	// Получаем информацию о выдаче книги по ID
	borrowByID, err := s.Repository.GetBorrowByID(borrowID)
	if err != nil {
		if errors.As(err, &ErrRecordNotFound) {
			return ErrBorrowNotFound
		}
		return err
	}

	// Проверяем, вернулась ли книга
	if borrowByID.ReturnDate != nil {
		//return fmt.Errorf("book already returned with ID %d", borrowID)
		return ErrBookAlreadyReturned
	}
	// Получаем информацию о книге по ID
	_, err = s.Repository.GetBookByBorrow(borrowID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return fmt.Errorf("book not found with borrow ID %d", borrowID)
		}
		return err
	}

	// Проверяем, существует ли пользователь
	//_, err = s.Repository.GetUserByID(borrowByID.UserID)
	//if err != nil {
	//	if errors.As(err, &ErrRecordNotFound) {
	//		return fmt.Errorf("user not found with ID %d", borrowByID.UserID)
	//	}
	//	return err
	//}
	if userID != borrowByID.UserID {
		return ErrDifferentUser
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
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("borrow not found with ID %d", borrowID)
		}
		return nil, err
	}

	return borrowByID, nil
}

func (s *Service) GetBorrowsByUser(userID int) ([]model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with ID %d", userID)
		}
		return nil, err
	}

	// Получаем все выдачи пользователя
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
	// Проверяем, существует ли книга
	_, err := s.Repository.GetBookByID(bookID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found with ID %d", bookID)
		}
		return nil, err
	}

	// Получаем все выдачи книги
	borrowsByBook, err := s.Repository.GetBorrowsByBook(bookID)
	if err != nil {
		return nil, err
	}
	if len(borrowsByBook) == 0 {
		return nil, fmt.Errorf("no borrows found with book ID %d", bookID)
	}

	return borrowsByBook, nil
}

func (s *Service) GetBorrowsByUserAndBook(userID, bookID int) ([]model.Borrow, error) {
	// Проверяем, существует ли пользователь
	_, err := s.Repository.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found with ID %d", userID)
		}
		return nil, err
	}

	// Проверяем, существует ли книга
	_, err = s.Repository.GetBookByID(bookID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("book not found with ID %d", bookID)
		}
		return nil, err
	}

	// Получаем все выдачи пользователя и книги
	borrowByUserAndBook, err := s.Repository.GetBorrowsByUserAndBook(userID, bookID)
	if err != nil {
		return nil, err
	}
	if len(borrowByUserAndBook) == 0 {
		return nil, fmt.Errorf("borrow not found with user ID %d and book ID %d", userID, bookID)
	}

	return borrowByUserAndBook, nil
}
