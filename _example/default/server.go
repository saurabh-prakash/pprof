package main

import (
	"github.com/saurabh-prakash/gin"
	"github.com/saurabh-prakash/pprof"
)

func main() {
	router := gin.Default()
	pprof.Register(router)
	router.Run(":8080")
}
