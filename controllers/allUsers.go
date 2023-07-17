package controllers

import (
	"net/http"
	"projetos-pessoais/crud-porfolio-backend/database"
	"projetos-pessoais/crud-porfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {
	var user []models.User
	database.DB.Find(&user)
	c.JSON(http.StatusOK, &user)
}
