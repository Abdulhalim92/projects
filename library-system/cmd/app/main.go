package main

import (
	"bufio"
	"fmt"
	"os"
	"projects/internal/book"
	"projects/internal/model"
	"projects/internal/user"
	"strings"
)

func main() {
	fmt.Println("Library System")

	// Инициализация книг
	//books := make(map[int]model.Book)
	//newBooks := book.NewBooks(books)

	jsonBooks := book.NewJSONBooks("books.json")
	bookService := book.NewService(*jsonBooks)

	// Инициализация пользователей
	jsonUsers := user.NewJSONUsers("users.json")
	userService := user.NewService(*jsonUsers)

	var optionsInput string

	var usrname string
	var password string
	var account *model.User

Register:
	for {
		fmt.Print("Options: signup, login\nExit: exit\n")
		fmt.Scan(&optionsInput)
		if optionsInput == "exit" {
			return
		}

		switch strings.ToLower(optionsInput) {
		case "login":
			if usrname == "exit" {
				return
			}

			fmt.Print("Username: ")
			fmt.Scan(&usrname)

			if usr, ok := userService.FindUserByUsername(usrname); ok {
				fmt.Print("Password: ")
				fmt.Scan(&password)
				for usr.Password != password {
					fmt.Println("Given password is not correct.")
					fmt.Println("To go back type 'back'")

					fmt.Print("Password: ")
					fmt.Scan(&password)
					if password == "back" {
						goto Register
					}
				}

				fmt.Println("Loged")
				account = usr
				break Register
			}

		case "signup":
			fmt.Print("Username: ")
			fmt.Scan(&usrname)

			fmt.Print("Password: ")
			fmt.Scan(&password)

			if usr, ok := userService.CreateUser(usrname, password); ok {
				account = usr
				fmt.Println("Signed up")
				break Register

			} else {
				fmt.Println("Error occured when creating user")
				return
			}
		}
	}

	dataInput := bufio.NewReader(os.Stdin)
	var title string
	var author string

	for {
		fmt.Println(account.Username)
		fmt.Print("Options: add, edit, find, delete\nExit: exit\n")
		fmt.Scan(&optionsInput)
		if optionsInput == "exit" {
			break
		}

		switch strings.ToLower(optionsInput) {
		case "add":

			fmt.Print("Type the book title\n")
			fmt.Print(">>> ")
			title, _ = dataInput.ReadString('\n')
			title = title[:len(title)-1]

			fmt.Print("Type the book's author\n")
			fmt.Print(">>> ")
			author, _ = dataInput.ReadString('\n')
			author = author[:len(author)-1]
			bookService.CreateBook(title, author)

		case "edit":
			var id int
			fmt.Print("Type the book id\n")
			fmt.Print(">>> ")
			fmt.Scan(&id)

			fmt.Print("Type new title\n")
			fmt.Print(">>> ")
			title, _ = dataInput.ReadString('\n')
			title = title[:len(title)-1]

			fmt.Print("Type new author\n")
			fmt.Print(">>> ")
			author, _ = dataInput.ReadString('\n')
			author = author[:len(author)-1]

			result := bookService.EditBook(id, title, author)
			if result {
				fmt.Printf("Book with id:%d  changed successfully\n", id)
			} else {
				fmt.Println("Error occured when editing the book")
			}

		case "find":
			var id int
			fmt.Print("Type the book id\n")
			fmt.Print(">>> ")
			fmt.Scan(&id)
			book, ok := bookService.FindBook(id)

			if ok {
				fmt.Println("Title: ", book.Title)
				fmt.Println("Author: ", book.Author)
			} else {
				fmt.Printf("Couldn't find the book with id:%d\n", id)
			}

		case "delete":
			var id int
			fmt.Print("Type the book id\n")
			fmt.Print(">>> ")
			fmt.Scan(&id)

			ok := bookService.RemoveBook(id)
			if ok {
				fmt.Printf("Book with id:%d deleted successfully \n", id)
			} else {
				fmt.Println("Error occured when deleting the book")
			}
		case "list":
			books := bookService.ListBooks()

			for _, v := range books {
				fmt.Printf("%d:\n  Title: %s\n  Author: %s\n", v.ID, v.Title, v.Author)
			}
		}
	}

	// // Имитируем создание книг
	// b1 := bookService.CreateBook("The Hobbit", "J.R.R Tolkien")
	// b2 := bookService.CreateBook("1984", "George Orwell")

	// // Имитируем создание пользователей
	// u1 := userService.CreateUser("johndoe", "password123")
	// u2 := userService.CreateUser("janedoe", "securepassword")

	// // Имитируем получение списка пользователей
	// listUsers := userService.ListUsers()
	// fmt.Println("Users in system:")
	// for _, u := range listUsers {
	// 	fmt.Printf("ID: %d, Username: %s, Password: %s\n", u.ID, u.Username, u.Password)
	// }

	// // Имитируем получение списка книг
	// listBooks := bookService.ListBooks()
	// fmt.Println("Books in library:")
	// for _, b := range listBooks {
	// 	fmt.Printf("ID: %d, Title: %s, Author: %s\n", b.ID, b.Title, b.Author)
	// }

	// // Имитируем обновление пользователя
	// updatedUser := userService.EditUser(u1.ID, "johnsmith", "newpassword")
	// if updatedUser {
	// 	fmt.Printf("User ID %d updated successfully\n", u1.ID)
	// }

	// // Имитируем обновление книг
	// updatedBook := bookService.EditBook(b1.ID, "The Hobbit: An Unexpected Journey", "J.R.R. Tolkien")
	// if updatedBook {
	// 	fmt.Printf("Book ID %d updated successfully: Title: %s, Author: %s\n", b1.ID, b1.Title, b1.Author)
	// }

	// // Имитируем удаление пользователя
	// deletedUser := userService.RemoveUser(u2.ID)
	// if deletedUser {
	// 	fmt.Printf("User ID %d deleted successfully\n", u2.ID)
	// }

	// // Имитируем удаление книги
	// deletedBook := bookService.RemoveBook(b2.ID)
	// if deletedBook {
	// 	fmt.Printf("Book ID %d deleted successfully\n", b2.ID)
	// }

	// // Имитируем получение пользователя по ID
	// foundUser := userService.FindUser(u1.ID)
	// if foundUser != nil {
	// 	fmt.Printf("Found User - ID: %d, Username: %s, Password: %s\n", foundUser.ID, foundUser.Username, foundUser.Password)
	// }

	// // Имитируем получение книги по ID
	// foundBook := bookService.FindBook(b1.ID)
	// if foundBook != nil {
	// 	fmt.Printf("Found Book - ID: %d, Title: %s, Author: %s\n", foundBook.ID, foundBook.Title, foundBook.Author)
	// }
}
