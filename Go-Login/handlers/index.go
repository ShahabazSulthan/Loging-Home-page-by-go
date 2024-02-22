package handlers

import (
	"fmt"
	"goLogin/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// Clear Cache
	ClearCache(c)

	// Checking Cookie Exits also Fetching details

	var name string
	cookie, err := c.Cookie("session")
	if err != nil {
		fmt.Println("Error in index rooter:", err)
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session, exist := data.Sessions[cookie]

	if exist {
		data, exist := data.UserData[session.Email]
		if exist {
			name = data.Name
			c.HTML(http.StatusOK, "index.html", name)
		}
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}
