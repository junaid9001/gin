package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("the secret key"))
	r.Use(sessions.Sessions("login", store))

	r.POST("/set", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(404, "error occured")
			return
		}

		if user.Name == "junaid" && user.Password == "1234" {
			session := sessions.Default(c)
			session.Set("username", user.Name)
			session.Save()
			c.String(200, "success")
			return

		} else {
			c.String(401, "invalid credential")
			return
		}

	})

	r.GET("/dashboard", func(c *gin.Context) {
		session := sessions.Default(c)
		val := session.Get("username")
		if val == nil {
			c.String(401, "no user found")
			return
		}
		c.JSON(200, gin.H{
			"username": val,
		})
	})

	r.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)

		session.Clear()
		session.Save()
		c.String(200, "logout confirmed")
	})
	r.Run(":8080")

}
