package user

import "projects/internal/model"

type Users struct {
	UserMap map[int]model.User
	LastId  int
}

func NewUsers(m map[int]model.User) *Users {
	return &Users{
		UserMap: m,
		LastId:  0,
	}
}
func (u *Users) AddUser(UserName string, password string) *model.User {
	u.LastId++
	user := model.User{
		UserId:   u.LastId,
		UserName: UserName,
		Password: password,
	}

	u.UserMap[u.LastId] = user

	us := u.UserMap[u.LastId]

	return &us
}
func (u Users) GetUsers() []model.User {
	x := make([]model.User, 0)
	for _, v := range u.UserMap {
		x = append(x, v)
	}
	return x
}
func (u *Users) GetUserById(id int) *model.User {
	for _, v := range u.UserMap {
		if v.UserId == id {
			return &v
		}
	}
	return nil
}
func (u *Users) DeleteUserById(id int) bool {
	for key, v := range u.UserMap {
		if v.UserId == id {
			delete(u.UserMap, key)
			return true
		}
	}
	return false
}
