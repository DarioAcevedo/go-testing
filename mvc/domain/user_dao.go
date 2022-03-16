package domain

import (
	"fmt"
	"net/http"

	"github.com/DarioAcevedo/go-testing/mvc/utils"
)

type userDao struct {}

type userDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

var(
	users = map[int64]*User{
		1: {
			Id: 1,
			FirstName: "Dario",
			LastName: "Acevedo",
			Email: "darioacevedo43@gmail.com",
		},
		2: {
			Id: 2,
			FirstName: "Alejandra",
			LastName: "Fernandez",
			Email: "afernandez@cornershopapp.com",
		},
		3: {
			Id: 3,
			FirstName: "Elisa",
			LastName: "Kauffmann",
			Email: "elisa.kauffmann@cornershopapp.com",
		},
		4: {
			Id: 4,
			FirstName: "Alejandro",
			LastName: "Sazo",
			Email: "alejandro.sazo@cornershopapp.com",
		},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = userDao{}
}


func (u userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	fmt.Println("Calling the database")
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}
	return user, nil
}