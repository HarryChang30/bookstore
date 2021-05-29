package app

import (
	"github.com/HarryChang30/bookstore/bookstore_users-api/controllers/ping"
	"github.com/HarryChang30/bookstore/bookstore_users-api/controllers/users"

	"github.com/gin-gonic/gin"
)

func StartApplication() {
	r := gin.Default()

	//routes
	r.GET("/health", ping.Ping)

	r.GET("/users/:user_id", users.GetUser)
	r.GET("/users/search", users.SearchUser)
	r.POST("/users", users.CreateUser)

	r.Run(":3000")
}
