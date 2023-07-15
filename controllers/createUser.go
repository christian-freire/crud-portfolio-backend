package controllers

import (
	"net/http"
	"projetos-pessoais/crud-porfolio-backend/database"
	"projetos-pessoais/crud-porfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
