package controllers

import (
	"net/http"
	"projetos-pessoais/crud-porfolio-backend/database"
	"projetos-pessoais/crud-porfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func DeleteUserByID(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")

	database.DB.Delete(&user, id)

	/* if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found:": "User not found in database."})
		return
	} */

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
