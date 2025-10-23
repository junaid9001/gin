package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	type User struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required,min=5"`
	}
	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"username": user.Name,
			"password": user.Password,
		})
	})
	r.Run(":8080")
}
