package main

import (
	"github.com/DenrianWeiss/barginerFish/web"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.LoadHTMLGlob("template/*")
	router.Static("/static", "./static")
	router.GET("/", web.RenderMainPage)
	router.POST("/add", web.AddItem)
	router.GET("/add", web.RenderCreatePage)
	router.GET("/items/:id", web.RenderItemPage)
	router.GET("/edit/:id", web.EditItem)
	router.POST("/edit/:id", web.EditItemApi)
	router.POST("/delete/:id/:token", web.DeleteItemApi)
	err := router.Run(":8080")
	if err != nil {
		return
	}
	select {}
}
