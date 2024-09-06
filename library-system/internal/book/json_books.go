package book

//
//import (
//	"fmt"
//	"projects/internal/model"
//	"projects/internal/utils"
//)
//
//const booksFile = "books.json"
//
//type JSONBooks struct {
//	filename string
//}
//
//func NewJSONBooks(filename string) *JSONBooks {
//	return &JSONBooks{
//		filename: filename,
//	}
//}
//
//func (b *JSONBooks) loadBooks() (map[int]model.Book, error) {
//	var books map[int]model.Book
//	err := utils.ReadJSONFromFile(b.filename, &books)
//	if err != nil {
//		return nil, err
//	}
//	return books, nil
//}
//
//func (b *JSONBooks) saveBooks(books map[int]model.Book) error {
//	return utils.WriteJSONToFile(b.filename, &books)
//}
//
//func (b *JSONBooks) AddBook(title, author string) model.Book {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return model.Book{}
//	}
//	lastID := 0
//
//	for id := range books {
//		if id > lastID {
//			lastID = id
//		}
//	}
//	lastID++
//
//	book := model.Book{
//		BookID: lastID,
//		Title:  title,
//		Author: author,
//	}
//
//	books[lastID] = book
//
//	err = b.saveBooks(books)
//	if err != nil {
//		fmt.Printf("Failed to save books: %v\n", err)
//		return model.Book{}
//	}
//
//	fmt.Printf("Book with tittle %s and author %s is created\n", book.Title, book.Author)
//
//	return book
//}
//
//func (b *JSONBooks) GetBooks() map[int]model.Book {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return nil
//	}
//
//	return books
//}
//
//func (b *JSONBooks) GetBookByID(id int) *model.Book {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return nil
//	}
//
//	book, exists := books[id]
//	if !exists {
//		return nil
//	}
//
//	return &book
//}
//
//func (b *JSONBooks) GetBooksByAuthor(author string) []model.Book {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return nil
//	}
//
//	var filteredBooks []model.Book
//
//	for _, book := range books {
//		if book.Author == author {
//			filteredBooks = append(filteredBooks, book)
//		}
//	}
//
//	return filteredBooks
//}
//
//func (b *JSONBooks) UpdateBook(id int, title, author string) bool {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return false
//	}
//
//	book, exists := books[id]
//	if !exists {
//		fmt.Printf("Book with id %d does not exist\n", id)
//		return false
//	}
//
//	book.Title = title
//	book.Author = author
//	books[id] = book
//	err = b.saveBooks(books)
//	if err != nil {
//		fmt.Printf("Failed to save books: %v\n", err)
//		return false
//	}
//
//	return true
//}
//
//func (b *JSONBooks) DeleteBook(id int) bool {
//	books, err := b.loadBooks()
//	if err != nil {
//		fmt.Printf("Failed to load books: %v\n", err)
//		return false
//	}
//
//	_, exists := books[id]
//	if !exists {
//		fmt.Printf("Book with id %d does not exist\n", id)
//		return false
//	}
//
//	delete(books, id)
//	err = b.saveBooks(books)
//	if err != nil {
//		fmt.Printf("Failed to save books: %v\n", err)
//		return false
//	}
//
//	return true
//}
