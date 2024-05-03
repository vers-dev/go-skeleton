package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	s := gin.Default()

	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := s.Run()

	if err != nil {
		panic(err)
	}
}
