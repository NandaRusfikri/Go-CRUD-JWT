package main

import (
	"./controllers"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	auth := middleware.Auth

	router := gin.Default()

	router.POST("/user", controllers.CreateUser)
	router.GET("/users", auth, controllers.GetUsers)
	router.PUT("/user", auth, controllers.UpdateUser)
	router.DELETE("/user", auth, controllers.DeleteUser)
	router.POST("/login", controllers.Login)

	router.Run(":3000")
}
