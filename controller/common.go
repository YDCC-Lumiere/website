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
	FirstName string
	LastName string
	Desc string
	Account string
}

func info(c *gin.Context) {
	view.Execute(c, "info", []InfoData{
		InfoData{
			Img: "static/images/nhu.png",
			FirstName: "Trần Nguyễn",
			LastName: "Huế Như",
			Desc: "Software Engineer",
			Account: "2power9",
		},
		InfoData{
			Img: "static/images/hai.png",
			FirstName: "Lê",
			LastName: "Đình Hải",
			Desc: "Software Engineer",
			Account: "haile01",
		},
		InfoData{
			Img: "static/images/khoa.png",
			FirstName: "Nguyễn Công",
			LastName: "Anh Khoa",
			Desc: "AI Engineer",
			Account: "ncakhoa",
		},
		InfoData{
			Img: "static/images/truc.png",
			FirstName: "Huỳnh Ngô",
			LastName: "Trung Trực",
			Desc: "AI Engineer",
			Account: "TrucJ",
		},
	})
}

func alert(c *gin.Context) {
	view.Execute(c, "alert-empty", nil)
}
