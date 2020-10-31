package app

import (
	"github.com/apm-dev/go-user-rest-api/controllers"
	"github.com/apm-dev/go-user-rest-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users", users.Index)
	router.GET("/users/:user_id", users.Show)
	router.POST("/users", users.Store)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
