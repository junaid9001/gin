package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Handleauth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		method := c.Request.URL.Path

		if method == "/dashboard" && username != "junaid" {
			c.Abort()
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("the secret key"))
	r.Use(sessions.Sessions("auth", store))
	r.Use(Handleauth())

	r.GET("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", "junaid")
		session.Save()
		c.JSON(200, gin.H{
			"session": "saved",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")

		c.JSON(200, gin.H{
			"username": username,
		})
	})

}
