package handlers

// fmt: This package implements formatted I/O with functions analogous to C's printf and scanf.
// goLogin/data: This likely refers to some custom data package in the goLogin project.
// github.com/gin-gonic/gin: Gin is a web framework for Go. This import is for handling HTTP requests and responses.
// github.com/google/uuid: This package provides a way to generate and manipulate UUIDs.

// A UUID – that's short for Universally Unique IDentifier,
// by the way – is a 36-character alphanumeric string that can be used to identify information.

import (
	"fmt"
	"goLogin/data"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.

// This function DeleteCookie is responsible for deleting a cookie named "session".

func DeleteCookie(c *gin.Context) {
	c.SetCookie("session", "", -1, "/", "", false, true)
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("Something Wrong in DeleteCookie Function", err)
	}
	delete(data.Sessions, cookie)
	fmt.Println("Cookie Deleted")
}

// This function CreateCookie is responsible for creating a new session cookie.

func CreateCookie(c *gin.Context) {
	fmt.Println("Cookie Created")
	id := uuid.NewString()
	email := c.PostForm("email")
	fmt.Println("Cookie id: ", id, "Email is: ", email)

	// Set Cookie and Session Datas

	session := data.SessionData{SessionId: id, Email: email}
	data.Sessions[id] = session

	c.SetCookie("session", id, 3600, "/", "", false, true)
	fmt.Println(data.Sessions)
}
