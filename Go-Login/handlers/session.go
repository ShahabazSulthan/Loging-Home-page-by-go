package handlers

import (
	"fmt"
	"goLogin/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

//This function VerifyCookie checks whether a session cookie exists and is valid.

func VerifyCookie(c *gin.Context) bool{

	cookie,err := c.Cookie("session")

	if err!=nil {
		fmt.Println("Error occurred while checking the session file in the VerifyCookie function:",err)
	}

	_,exist := data.Sessions[cookie]

	if exist {
		c.Redirect(http.StatusFound,"/index")
		return true
	}
	return false
}
