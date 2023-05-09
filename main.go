package main

import (
	"sin"
	"fmt"
	"net/http"

)

func main() {
	r := sin.New()
	
	r.GET("/",func(c *sin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Cui Taiyu</h1>")
	})

	r.GET("/hello",func(c *sin.Context) {
		c.String(http.StatusOK,"hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *sin.Context) {
		c.JSON(http.StatusOK, sin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	fmt.Printf("Run success")
	r.Run(":8080")
}
