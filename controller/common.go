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
}

func info(c *gin.Context) {
	view.Execute(c, "info", []InfoData{
		InfoData{
			Img: "",
			FirstName: "Trần Nguyễn",
			LastName: "Huế Như",
			Desc: "Software Engineer",
		},
		InfoData{
			Img: "",
			FirstName: "Lê",
			LastName: "Đình Hải",
			Desc: "Software Engineer",
		},
		InfoData{
			Img: "https://media.discordapp.net/attachments/1221316046171017239/1228657723986546698/e91124a81138d6668f29.jpg?ex=662cd7ac&is=661a62ac&hm=99e69276a29f007445e7b86dc6fdd514815b374b530c3ab7bc0c66a86b641569&=&format=webp",
			FirstName: "Nguyễn Công",
			LastName: "Anh Khoa",
			Desc: "AI Engineer",
		},
		InfoData{
			Img: "https://cdn.discordapp.com/attachments/1221316046611284009/1228656737813270591/IMG_20231006_095308.jpg?ex=662cd6c1&is=661a61c1&hm=88e1f55b53d76fa48ca9e3c4e04142cbace566740cbed0e2c33c10e33967bc6b&",
			FirstName: "Huỳnh Ngô",
			LastName: "Trung Trực",
			Desc: "AI Engineer",
		},
	})
}

func alert(c *gin.Context) {
	view.Execute(c, "alert-empty", nil)
}
