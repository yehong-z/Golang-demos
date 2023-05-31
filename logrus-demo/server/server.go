package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyMiddleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("MyMiddleware1: before request")
		c.String(http.StatusOK, "Hello, World!")
		go c.Next()
		c.Abort()
		fmt.Println("MyMiddleware1: after request")
	}
}

func MyMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(time.Second * 10)
		fmt.Println("MyMiddleware2: before request")
		c.Next()
		fmt.Println("MyMiddleware2: after request")
	}
}

func Run() {
	r := gin.Default()

	r.Use(MyMiddleware1())
	r.Use(MyMiddleware2())

	r.GET("/hello", func(c *gin.Context) {
		fmt.Println("Handling request")
		// c.String(http.StatusOK, "Hello, World!")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
