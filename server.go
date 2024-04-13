package main

import (
  "github.com/gin-gonic/gin"

  "go-mine/controller"
  "go-mine/middleware"
  "go-mine/view"
)

func main() {

  r := gin.Default()
  r.Static("/static", "./static")
  r.Use(gin.Logger())
  r.Use(middleware.TextHTML())
  view.Initialize("templates/*")

  controller.DocsController(r.Group("docs"))
  r.Run()
}
