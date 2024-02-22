package handlers

import (
	"goLogin/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	ClearCache(c)

	//Check Cookie Exit
	cookieExit := VerifyCookie(c)
	if !cookieExit {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func PostLogin(c *gin.Context) {
	ClearCache(c)

	// Data Fetching From Form
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Compare Previous data and Cookie Settiings

	val, exist := data.UserData[email]
	if exist && val.Email == email && val.Password == password {
		CreateCookie(c)
		c.Redirect(http.StatusFound, "/index")
		return
	} else {
		c.HTML(http.StatusOK, "login.html", "Enter Valid Data")
	}
}
