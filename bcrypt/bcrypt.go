package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := gin.Default()

	type User struct {
		Password string `json:"password"`
	}

	r.POST("/signup", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": "some internal error",
			})
			return
		}
		passtohash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "internal error",
			})
			return
		}
		c.JSON(200, gin.H{
			"salt": string(passtohash),
		})
	})

	r.Run(":8080")
}
