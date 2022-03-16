package controllers

import (
	"fmt"
	"net/http/httptest"
	"testing"

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
	w := httptest.NewRecorder()
	for _, tc := range uCTests {
		CallGetUserTest(tc.n, w)
		assert.Equal(t, tc.exp, w.Body.String())
	}
}


func CallGetUserTest(input int, w *httptest.ResponseRecorder) {
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key:"user_id", Value: fmt.Sprint(input)}}
	GetUser(c)
}