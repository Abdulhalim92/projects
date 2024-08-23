package main

import (
	"fmt"
	"projects/internal/book"
)

func main() {

	x := book.NewJSONBooks("file.json")
	s := book.NewService(x)
	for {
		var operation string
		fmt.Println("Choose operation:")
		fmt.Scan(&operation)
		if operation == "exit" {
			return
		}
		switch operation {
		case "getbooks":
			books := s.ListBooks()
			for _, book := range books {
				fmt.Println(book)
			}
		case "getbookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			book := s.FindBook(id)
			fmt.Println(book)
		case "deletebookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			b := s.RemoveBook(id)
			if !b {
				fmt.Println("Error while deleting")
			}
			fmt.Println("Book deleted")
		case "changebookbyid":
			var id int
			fmt.Println("Enter id:")
			fmt.Scan(&id)
			var author string
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			var title string
			fmt.Println("Enter title:")
			fmt.Scan(&title)
			b := s.EditBook(id, title, author)
			if !b {
				fmt.Println("Error while editing a book")
			}
			fmt.Println("Book edited")
		case "getbooksbyauthor":
			var author string
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			books := s.FindBooksByAuthor(author)
			for _, book := range books {
				fmt.Println(book)
			}
		case "addbook":
			var author, title string
			fmt.Println("Enter author:")
			fmt.Scan(&author)
			fmt.Println("Enter title:")
			fmt.Scan(&title)
			s.CreateBook(title, author)
		}
	}
}
