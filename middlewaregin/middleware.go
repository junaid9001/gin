package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		path := c.Request.URL.Path

		if path == "/dashboard" {
			get := session.Get("username")
			if get != "junaid" {
				c.Abort()
				log.Print("wrong user tried to access dashboard")
			} else {
				c.Next()
				log.Print("succesfully opened dashboard")
			}

		} else {

			c.Next()
		}
	}
}

func main() {

	r := gin.Default()

	username := "junaid"
	password := "1234"

	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	store := cookie.NewStore([]byte("the secret key"))

	r.Use(sessions.Sessions("auth", store))

	r.Use(Auth())

	r.POST("/log", func(c *gin.Context) {
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"status": "somethin gone wrong",
			})
			return
		}

		session := sessions.Default(c)

		session.Set("username", user.Name)
		session.Set("password", user.Password)
		session.Save()
		c.JSON(200, gin.H{
			"status": "session created succefully",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		session := sessions.Default(c)
		Username := session.Get("username")
		Password := session.Get("password")

		if Username != username && Password != password {
			c.JSON(400, gin.H{
				"status": "bad response",
			})
			return
		}

		c.JSON(200, gin.H{
			"username": Username,
		})
	})

	r.GET("logout", func(c *gin.Context) {
		session := sessions.Default(c)

		session.Clear()
		session.Save()

	})

	r.Run(":8081")

}
