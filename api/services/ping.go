package services

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) (interface{}, error) {
	return "PONG", nil
}
