package controller

import (
  "github.com/gin-gonic/gin"
	"go-mine/view"
)

func CommonController(g *gin.RouterGroup) {
	// Conventional

	// HTMX
	g.POST("/alert", alert)
}

func alert(c *gin.Context) {
	view.Execute(c, "alert-empty.tmpl", nil)
}
