package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	//simple route
	route.GET("/", webRoot)

	//route with param
	route.GET("param/:name", paramRoute)

	//route group
	v1 := route.Group("group")
	{
		v1.GET("/first", firstGroup)
		v1.GET("/second", secondGroup)
	}

	//route with middleware
	route.GET("/middle", firstMiddleware, secondMiddleware,handler)

	//route with group & middleware
	group2 := route.Group("/group2", firstMiddleware)
	group2.GET("first", )

	//post get form
	route.POST("/post", postParam)

	route.Run(":9205")
}

func webRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}

func paramRoute(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, "Input Param : %s", name)
}

func firstGroup(context *gin.Context) {
	context.String(http.StatusOK, "You Got First Group Route")
}

func secondGroup(context *gin.Context) {
	context.String(http.StatusOK, "You Got Second Group Route")
}

func handler(c *gin.Context) {
	log.Println("running At Handler")
}

func firstMiddleware(c *gin.Context) {
	log.Println("running At   First   Middleware")
	c.Next()
	log.Println("running AFter First   Middleware")
}

func secondMiddleware(c *gin.Context) {
	log.Println("running At  Second   Middleware")
	c.Next()
	log.Println("running AFter  Second   Middleware")
}

func postParam(c *gin.Context) {
	param := c.PostForm("param1")
	paramDefault := c.DefaultPostForm("param2", "Default")

	c.JSON(200, gin.H{
		"Function" : "post",
		"param1" : param,
		"paramDefault" : paramDefault,
	})
}
