package user

import "projects/library-system/internal/model"

// import "projects/internal/model"

type Service struct {
	users Users
}

func NewService(users Users) *Service {
	return &Service{
		users: users,
	}
}

func (s *Service) CreateUser(username, password string) model.User {
	return s.users.AddUser(username, password)
}

func (s *Service) ListUsers() []model.User {
	return s.users.GetUsers()
}

func (s *Service) FindUser(id int) *model.User {
	return s.users.GetUserByID(id)
}

func (s *Service) EditUser(id int, username, password string) bool {
	return s.users.UpdateUser(id, username, password)
}

func (s *Service) RemoveUser(id int) bool {
	return s.users.DeleteUser(id)
}