package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		for {
			fmt.Println("a")
		}
	})
	r.Run("127.0.0.1:9995") // listen and serve on 0.0.0.0:8080
}
