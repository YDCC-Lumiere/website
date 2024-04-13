package controller

import (
	"log"
  "github.com/gin-gonic/gin"
	"go-mine/view"
)

func DocsController(g *gin.RouterGroup) {
	// Conventional
	g.GET("/", index)

	// HTMX
}

func index(c *gin.Context) {
	log.Println("/docs")
	view.Alert("success", "Test", "testtttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttttt")
	view.Execute(c, "docs.tmpl", nil)
}
