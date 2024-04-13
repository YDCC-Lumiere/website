package controller

import (
  "github.com/gin-gonic/gin"
	"go-mine/view"
)

func CommonController(g *gin.RouterGroup) {
	// Conventional
	g.GET("/", index)
	// g.GET("/getting-started", tutorial)
	g.GET("/documentation", documentation)
	// g.GET("/about-us", info)

	// HTMX
	g.POST("/alert", alert)
}

func index(c *gin.Context) {
	view.Execute(c, "index", nil)
}

// func tutorial(c *gin.Context) {
// 	view.Execute(c, "tutorial", nil)
// }

func documentation(c *gin.Context) {

	view.Execute(c, "documentation", nil)
}

// func info(c *gin.Context) {
// 	view.Execute(c, "info", nil)
// }

func alert(c *gin.Context) {
	view.Execute(c, "alert-empty", nil)
}
