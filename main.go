package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", WebRoot)

	v1 := route.Group("api/v1")
	{
		v1.GET("login", login)
	}

	route.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	route.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	route.Run(":9205")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}

func login(context *gin.Context) {
	context.String(http.StatusOK, "login route")
}
