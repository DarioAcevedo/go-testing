package app

import(
	"github.com/DarioAcevedo/go-testing/mvc/controllers"
) 


func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}