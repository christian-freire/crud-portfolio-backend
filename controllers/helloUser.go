package controllers

import "github.com/gin-gonic/gin"

func HelloUser(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says:": "Hello " + name + ", welcome to this API!",
	})

}
