package main

import (
	"fmt"
	"log"
	"projects/internal/user"
)

func main() {
	j := user.NewJSONUsers("sst.json")
	s := user.NewService(j)
	x := s.ListUsers()
	fmt.Println(x)
	fmt.Println(s.ListUserById(2))
	b := s.RemoveUserById(5)
	if !b {
		log.Fatal("Ошибка!!!")
	}
}
