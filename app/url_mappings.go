package app

import (
	"github.com/apm-dev/go-user-rest-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
