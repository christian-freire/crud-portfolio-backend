package routes

import (
	"projetos-pessoais/crud-porfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/users", controllers.AllUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.GET("/:name", controllers.HelloUser)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUserByID)
	r.Run()
}
