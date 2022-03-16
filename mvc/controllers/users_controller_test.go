package controllers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/DarioAcevedo/go-testing/mvc/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


var uCTests = []struct {
	n int // input
	exp string // expected value
} {
	{1, "Dario"},
	{2, "Alejandra"},
	{3, "Elisa"},
	{4, "Alejandro"},
}



func TestUserController (t *testing.T) {
	gin.SetMode(gin.TestMode)
	for _, tc := range uCTests {

		name := fmt.Sprintf("Test with user: %v", tc.exp)

		t.Run(name, func(t *testing.T) {
			
			t.Parallel()

			w := httptest.NewRecorder()
			CallGetUserTest(tc.n, w)
			var usr domain.User 
			err := json.Unmarshal(w.Body.Bytes(), &usr)
			if err != nil {
				t.Errorf("Got an error unmarshalling: %v", err)
			}
			assert.Equal(t, tc.exp, usr.FirstName)
		})

	}
}


func CallGetUserTest(input int, w *httptest.ResponseRecorder) {
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key:"user_id", Value: fmt.Sprint(input)}}
	GetUser(c)
}