package user

import (
	"quizgo/internal/model"
	"sync"
)

type Users struct {
	l        sync.RWMutex
	UsersMap map[int]model.User
}

func NewUsers() *Users {
	return &Users{
		UsersMap: map[int]model.User{},
	}
}

func (u *Users) Add(login, pswd string) model.User {
	u.l.Lock()
	defer u.l.Unlock()
	lastID := len(u.UsersMap)
	user := model.User{
		ID:       lastID,
		Login:    login,
		Password: pswd,
	}

	u.UsersMap[lastID] = user

	return user
}

func (u *Users) GetUsers() map[int]model.User {
	u.l.RLock()
	defer u.l.RUnlock()
	return u.UsersMap
}

func (u *Users) GetUserByID(id int) (model.User, bool) {
	u.l.RLock()
	defer u.l.RUnlock()
	user, ok := u.UsersMap[id]
	return user, ok
}

func (u *Users) GetUserByLogin(login string) (model.User, bool) {
	u.l.RLock()
	defer u.l.RUnlock()

	for k := range u.UsersMap {
		if u.UsersMap[k].Login == login {
			return u.UsersMap[k], true
		}
	}
	return model.User{}, false
}

func (u *Users) UpdateUser(user model.User) bool {
	for k := range u.UsersMap {
		if k == user.ID {
			u.UsersMap[k] = user
			return true
		}
	}
	return false
}

func (u *Users) DeleteUser(id int) bool {
	for k := range u.UsersMap {
		if k == id {
			delete(u.UsersMap, id)
			return true
		}
	}
	return false
}
