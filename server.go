package main

import (
  "github.com/gin-gonic/gin"

  "go-mine/controller"
  "go-mine/middleware"
  "go-mine/view"
  "go-mine/config"
)

func main() {
  //TODO: fix this
  config.Init("development")

  r := gin.Default()
  r.Static("/static", "./static")
  r.Use(gin.Logger())
  r.Use(middleware.TextHTML())
  view.Initialize("templates/*")

  controller.CommonController(r.Group("/"))
  r.Run()
}
