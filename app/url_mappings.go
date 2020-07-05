package app

import (
	"github.com/LonnieCoffman/bookstore_users-api/controllers/ping"
	"github.com/LonnieCoffman/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/users/:user_id", users.GetUser)
	// }
}
