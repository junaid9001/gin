package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secretkey"))

	r.Use(sessions.Sessions("hash", store))
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	r.POST("/sign", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		session := sessions.Default(c)

		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		session.Set("hashedpass", hashed)
		session.Save()
		c.JSON(200, gin.H{
			"status": "pass hashed succesfully",
		})

	})

	r.GET("/log", func(c *gin.Context) {
		session := sessions.Default(c)

		pass := session.Get("hashedpass")

		c.JSON(200, gin.H{
			"hashed pass got ": pass,
		})

	})
	r.Run(":8080")
}
