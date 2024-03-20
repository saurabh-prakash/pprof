package main

import (
	"net/http"

	"github.com/saurabh-prakash/gin"
	"github.com/saurabh-prakash/pprof"
)

func main() {
	router := gin.Default()
	adminGroup := router.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "foobar" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")
	router.Run(":8080")
}
