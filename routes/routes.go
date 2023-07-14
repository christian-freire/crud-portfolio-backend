package routes

import (
	"projetos-pessoais/crud-porfolio-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/user", controllers.AllUsers)
	r.GET("/:name", controllers.HelloUser)
	r.Run()
}
