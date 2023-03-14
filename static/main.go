package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	// Serve index.html from the directory containing main.go
	r.StaticFile("/", "./index.html")
	r.GET("/hello", guidMiddleware())
	r.GET("/world", guidMiddleware())

	r.Run(":8080")
}

// add the middleware function
func guidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "요청 경로??")

		switch c.Request.URL.Path {
		case "/hello":
			uuid := uuid.New()
			c.Set("uuidHELLO", uuid)
			break
		case "/world":
			uuid := uuid.New()
			c.Set("uuidWORLD", uuid)
			break
		}

		c.Next()
	}
}

func Resp1(c *gin.Context) {
	val, boo := c.Get("uuidHELLO")
	if !boo {
		c.JSON(400, "잘못된 요청")
	}

	c.JSON(200, "WWWWWWWworld"+fmt.Sprintf("앞에서 받아온 값은: %s", val))
}

func Resp2(c *gin.Context) {
	val, boo := c.Get("uuidWORLD")
	if !boo {
		c.JSON(400, "잘못된 요청")
	}

	c.JSON(200, "hellooOOOOOOO "+fmt.Sprintf("앞에서 받아온 값은: %s", val))
}
