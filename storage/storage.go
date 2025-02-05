package storage

import (
	"estudos/golang/model"
	"sync"
)

// Variáveis auxiliares
var (
	users  []model.User
	nextId int = 1
	mutex  sync.Mutex
)

// Listar usuário
func ListUsers() []model.User {
	mutex.Lock()
	defer mutex.Unlock()
	return users
}

// Criar usuário
func CreateUser(new model.User) model.User {
	mutex.Lock()
	defer mutex.Unlock()
	new.ID = nextId

	users = append(users, new)
	nextId++
	return new
}

// Obtém usuário pelo ID
func GetUser(id int) (model.User, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return model.User{}, false
}
