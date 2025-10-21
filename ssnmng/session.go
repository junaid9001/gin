package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("the secret key"))
	r.Use(sessions.Sessions("cookie", store))

	r.GET("/test", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("name", "junaid")

		if err := session.Save(); err != nil {
			c.JSON(500, gin.H{
				"status": "fail",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	r.GET("/getcookie", func(c *gin.Context) {
		session := sessions.Default(c)
		val := session.Get("name")
		if valstr, ok := val.(string); ok {
			c.String(200, "got it "+valstr)
		} else {
			c.String(404, "not found")
		}
	})

	r.GET("/del", func(c *gin.Context) {
		session := sessions.Default(c)

		session.Delete("name")
		session.Save()

		c.String(200, "name is deleted")
	})

	r.Run(":8080")

}
