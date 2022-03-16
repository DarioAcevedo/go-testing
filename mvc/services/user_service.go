package services

import (
	"github.com/DarioAcevedo/go-testing/mvc/domain"
	"github.com/DarioAcevedo/go-testing/mvc/utils"
) 

type userService struct {}

var(
	UserService userService
)

func (u userService)GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}