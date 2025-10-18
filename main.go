package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	var user User

	r.POST("/", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&user); err != nil {
			return
		}

		c.JSON(200, gin.H{
			"name": user.Name,
			"age":  user.Age,
		})
		fmt.Println("name is ", user.Name, " age is", user.Age)

	})

	api := r.Group("api")

	api.GET("first", func(c *gin.Context) {
		c.String(200, "on first")
	})

	api.GET("second", func(c *gin.Context) {
		c.String(200, "on second")
	})

	r.GET("dynamic/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"dynamicpart": id,
		})
	})

	r.Run(":8080")
}
