package controllers

import (
	"net/http"
	"projetos-pessoais/crud-porfolio-backend/database"
	"projetos-pessoais/crud-porfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {

	var user models.User
	id := c.Params.ByName("id")

	database.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found:": "User not found in database."})
		return
	}

	c.JSON(http.StatusOK, user)
}
