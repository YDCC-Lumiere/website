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
	g.GET("/about-us", info)

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

type InfoData struct {
	Img string
	Name string
	Desc string
}

func info(c *gin.Context) {
	view.Execute(c, "info", []InfoData{
		InfoData{
			Img: "",
			Name: "Trần Nguyễn Huế Như",
			Desc: "...",
		},
		InfoData{
			Img: "",
			Name: "Lê Đình Hải",
			Desc: "...",
		},
		InfoData{
			Img: "",
			Name: "Nguyễn Công Anh Khoa",
			Desc: "...",
		},
		InfoData{
			Img: "",
			Name: "Huỳnh Ngô Trung Trực",
			Desc: "...",
		},
	})
}

func alert(c *gin.Context) {
	view.Execute(c, "alert-empty", nil)
}
