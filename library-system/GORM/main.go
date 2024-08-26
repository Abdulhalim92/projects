package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Creating Database connection
	dsn := "host=localhost user=humo password=humo dbname=Humo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connecting to database: %v", err)
	}
	// 1
	//SelectUsers(db)
	// 2
	//SelectBooks(db)
	// 3
	//SelectAuthors(db)
	// 4
	//GetBooksOnAuthorName(db)
	// 5
	//SelectProfilesWithEmail(db)
	// 6
	//UpdatePasswordOfAlice(db)
	// 7
	//CountBooks(db)
	// 8
	//GetBorrowedBooks(db)
	// 9
	//ShowProfileOfAllUsers(db)
	// 10

	// 11
	//UpdateAddressOfBob(db)
	// 12
	//ShowAuthorOfWarAndPeace(db)
	// 13
	//ShowAmountOfBooksOfEachAuthor(db)
	// 14
	//SelectBooksAfterDate(db)
	// 15
	// ShowAllNoReturnbooks(db)
	// 16
	//ShowAllUsersWithoutProfiles(db)
	// 17
	//
	// 18
	//UpdateBiographyOfLeo(db)
	// 19
	ShowUsersWithMoretakingbooksthanone(db)
}
