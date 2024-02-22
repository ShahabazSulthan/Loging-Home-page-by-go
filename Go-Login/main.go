package main

import (
	"fmt"
	"goLogin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	// This line initializes a new instance of the Gin router with the default middleware stack.
	r := gin.Default()

	// This line tells Gin to load HTML templates from the "view" directory with the "*.html" file extension.
	// These templates can be rendered in response to HTTP requests.
	r.LoadHTMLGlob("view/*.html")

	//This line serves static files (e.g., CSS, JavaScript) located in the "static" directory under the "/static" route.
	// This allows the web server to serve these files directly to clients.
	r.Static("/static", "./static")

	//These lines define routes and associate them with specific handler functions from the handlers package.
	r.GET("/login", handlers.Login)
	r.POST("/login", handlers.PostLogin)
	r.GET("/index", handlers.Index)
	r.GET("/logout", handlers.Logout)

	err := r.Run(":7000")

	if err != nil {
		fmt.Println("Error in server start,", err)
	}
}
