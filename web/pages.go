package web

import (
	"github.com/DenrianWeiss/barginerFish/service/db"
	"github.com/DenrianWeiss/barginerFish/service/render"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

func RenderMainPage(ctx *gin.Context) {
	provider := db.GetDb()
	items, _ := provider.GetRecentItems()
	// Render items
	rendered := ""
	for _, item := range items {
		rendered += string(render.RenderCard(render.ConvertItemToRenderCard(item)))
	}
	ctx.HTML(200, "index.html", gin.H{
		"title": "Catball Shop",
		"items": template.HTML(rendered),
	})
}

func RenderCreatePage(ctx *gin.Context) {
	ctx.HTML(200, "add.html", gin.H{})
}

func RenderItemPage(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	item, err := db.GetDb().GetItemById(uint(idInt))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.HTML(200, "item.html", gin.H{
		"name":  item.ItemName,
		"price": item.ItemPrice,
		"desc":  template.HTML(render.RenderRichText([]byte(item.ItemDesc))),
		"full":  template.HTML(render.RenderRichText([]byte(item.ItemFullDesc))),
		"image": item.ItemImage,
	})
}

func EditItem(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	item, err := db.GetDb().GetItemById(uint(idInt))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.HTML(200, "edit.html", gin.H{
		"name":  item.ItemName,
		"price": item.ItemPrice,
		"desc":  item.ItemDesc,
		"full":  item.ItemFullDesc,
		"image": item.ItemImage,
		"id":    id,
	})
}
