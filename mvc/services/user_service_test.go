package services

import (
	"net/http"
	"testing"

	"github.com/DarioAcevedo/go-testing/mvc/domain"
	"github.com/DarioAcevedo/go-testing/mvc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testifyMockUserDao struct {
	mock.Mock
}

func (m *testifyMockUserDao) GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	args := m.Called()
	var user *domain.User
	var err *utils.ApplicationError
	if args[0] != nil {
		user = args.Get(0).(*domain.User)
	}
	if args[1] != nil {
		err = args.Get(1).(*utils.ApplicationError)
	}
	return user, err
}

type mockUserDao struct {}

func (u mockUserDao) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserMock(userId)
}

var( 
	mockedUserDao mockUserDao
	getUserMock func (int64) (*domain.User, *utils.ApplicationError)
)

 func init () {
 	domain.UserDao = mockedUserDao
 }
func TestGetUserServiceUserNotFound (t *testing.T) {

	getUserMock = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
	}

	user, err := UserService.GetUser(5)
	assert.Nil(t, user, "user should be nil")
	assert.NotNil(t, err, "err should not be nil")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "status code should be 404")
	assert.EqualValues(t, "User not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}
func TestGetUserService(t *testing.T) {
	mockedObj := new(testifyMockUserDao)
	mockedObj.On("GetUser").Return(&domain.User{
		Id: 1,
		FirstName: "Dario",
		LastName: "Acevedo",
		Email: "dariotest@gmail.com",
	}, nil)

	user, err := UserService.GetUser(1)

	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, user, "user should not be nil")
	assert.EqualValues(t, 1, user.Id)

}
