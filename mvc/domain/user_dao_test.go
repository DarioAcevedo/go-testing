package domain

import (
	"net/http"
	"testing"
)

//Test if id not in database
func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(5)
	if user != nil {
		t.Error("User should be nil")
	}
	if err == nil {
		t.Error("err should not be nil")
	} else {
		if sc := err.StatusCode; sc != http.StatusNotFound {
			t.Errorf("status should be 404, got %d", sc)
		}
		if got := err.Message; got != "User not found" {
			t.Errorf("expected message 'User not found' but got %s", got)
		}
		if code := err.Code; code != "not_found" {
			t.Errorf("expected error code 'not_found' but got %s", code)
		}
	}
}

func TestGetUserFound(t *testing.T) {
	user, err := UserDao.GetUser(1)

	if user == nil {
		t.Error("user should not be nil")
	} else {
		if id := user.Id; id != 1 {
			t.Errorf("expected id to be 1 but got %d", id)
		}
	}
	if err != nil {
		t.Error("err should be nil")
	}

}