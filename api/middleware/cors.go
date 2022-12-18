package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HelloWorld(c *gin.Context) {

	log.Println("Message test middleware")

	c.Next()
}
