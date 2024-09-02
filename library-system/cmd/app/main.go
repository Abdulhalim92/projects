package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	/*
		db, err := ConnectToDb()
		if err != nil {
			log.Fatal(err)
		}
		UserRep := UserDataBase.NewUserRepository(db)
		UserService := UserDataBase.NewService(UserRep)
		BookRep := BookDataBase.NewBookRepository(db)
		BookService := BookDataBase.NewService(BookRep)
		for {
			var object string
			fmt.Println("Choose an object: ")
			fmt.Scan(&object)
			if object == "user" {
				var operation string
				fmt.Println("Choose an operation: ")
				fmt.Scan(&operation)
				if operation == "add" {
					var user model.User
					fmt.Println("Username:")
					fmt.Scan(&user.Username)
					fmt.Println("Password:")
					fmt.Scan(&user.Password)
					u, err := UserService.CreateUser(user)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Printf("%v successfully created\n", *u)
					}
				} else if operation == "delete" {
					var id int
					fmt.Println("Id of the user: ")
					fmt.Scan(&id)
					b, err := UserService.RemoveUser(id)
					if !b {
						log.Println(err)
					} else {
						fmt.Printf("User with id %d successfully deleted\n", id)
					}
				} else if operation == "edit" {
					var user model.User
					fmt.Println("Id of the user: ")
					fmt.Scan(&user.Userid)
					fmt.Println("New Username: ")
					fmt.Scan(&user.Username)
					fmt.Println("New Password: ")
					fmt.Scan(&user.Password)
					b, err := UserService.EditUser(user)
					if !b {
						log.Println(err)
					} else {
						fmt.Printf("User with id %d successfully edited\n", user.Userid)
					}
				} else if operation == "getall" {
					users, err := UserService.ListUsers()
					if err != nil {
						log.Println(err)
					} else {
						for _, user := range users {
							fmt.Printf("User: %d %s %s\n", user.Userid, user.Username, user.Password)
						}
					}
				} else if operation == "getuser" {
					var id int
					fmt.Println("ID of the user: ")
					fmt.Scan(&id)
					user, err := UserService.ListUserById(id)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Printf("User: %d %s %s\n", user.Userid, user.Username, user.Password)
					}
				} else {
					fmt.Println("Such operation doesn't exist!!!")
				}
			} else if object == "book" {
				var operation string
				fmt.Println("choose an operation: ")
				fmt.Scan(&operation)
				if operation == "getall" {
					books, err := BookService.ListBooks()
					if err != nil {
						log.Println(err)
					} else {
						for _, book := range books {
							fmt.Printf("Book: %d %s %d\n", book.Bookid, book.Title, book.Authorid)
						}
					}
				} else if operation == "add" {
					var book model.Book
					fmt.Println("Title: ")
					fmt.Scan(&book.Title)
					fmt.Println("Its AuthorID: ")
					fmt.Scan(&book.Authorid)
					b, err := BookService.CreateBook(book)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Printf("Book %v successfully created\n", *b)
					}
				} else if operation == "getbook" {
					var id int
					fmt.Println("ID: ")
					fmt.Scan(&id)
					book, err := BookService.FindBook(id)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Printf("Book: %d %s %d\n", book.Bookid, book.Title, book.Authorid)
					}
				} else if operation == "getallbooksbyauthor" {
					var AuthorID int
					fmt.Println("AuthorID: ")
					fmt.Scan(&AuthorID)
					books, err := BookService.FindBooksByAuthor(AuthorID)
					if err != nil {
						log.Println(err)
					} else {
						for _, book := range books {
							fmt.Printf("Book: %d %s %d\n", book.Bookid, book.Title, book.Authorid)
						}
					}
				} else if operation == "delete" {
					var id int
					fmt.Println("ID: ")
					fmt.Scan(&id)
					b, err := BookService.RemoveBook(id)
					if !b {
						log.Println(err)
					} else {
						fmt.Printf("Book with id %d successfully deleted\n", id)
					}
				} else if operation == "edit" {
					var book model.Book
					fmt.Println("ID: ")
					fmt.Scan(&book.Bookid)
					fmt.Println("NewTitle: ")
					fmt.Scan(&book.Title)
					fmt.Println("NewAuthorID: ")
					fmt.Scan(&book.Bookid)
					b, err := BookService.EditBook(book)
					if !b {
						log.Println(err)
					} else {
						fmt.Printf("Book with id %d successfully edited\n", book.Bookid)
					}
				} else {
					fmt.Println("Such operation doesn't exist!!!")
				}
			} else if object == "exit" {
				return
			} else {
				fmt.Println("no such object")
			}
		}
	*/
}
func ConnectToDb() (*gorm.DB, error) {
	dsn := "host=localhost user=humo password=humo dbname=Humo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return nil, err
	}
	return db, nil
}
