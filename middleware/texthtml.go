package middleware

import (
  "github.com/gin-gonic/gin"
)

func TextHTML() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()
    c.Header("Content-Type", "text/html charset=utf-8")
  }
}
