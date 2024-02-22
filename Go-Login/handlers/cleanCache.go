package handlers

// "net/http" provides HTTP client and server implementations.
// "time" provides functionality for handling time-related operations.
// "github.com/gin-gonic/gin" is a popular web framework for building RESTful APIs in Go.

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// This line defines a function named ClearCache which takes a parameter c of type *gin.Context.
//  This function is intended to clear the cache related headers in an HTTP response.

func ClearCache(c *gin.Context) {

	// Cache controlls

	c.Header("Cache-Control", "no-store,no-cache,must-revalidate,max-age=0")
	c.Header("Expires", time.Time{}.UTC().Format(http.TimeFormat))

	// c.Header("Cache-Control", "no-cache, no-store, no-transform, must-revalidate, private, max-age=0")
	// c.Header("Pragma", "no-cache")
	// c.Header("X-Accel-Expires", "0")
}

/*the ClearCache function is intended to be used as a middleware in an HTTP request handling chain to ensure
that the response is not cached by setting appropriate cache control headers.*/
