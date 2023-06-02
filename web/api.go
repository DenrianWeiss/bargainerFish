package web

import (
	"github.com/DenrianWeiss/barginerFish/model"
	"github.com/DenrianWeiss/barginerFish/service/db"
	"github.com/DenrianWeiss/barginerFish/service/env"
	"github.com/DenrianWeiss/barginerFish/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strconv"
	"strings"
)

func AddItem(ctx *gin.Context) {
	imageFile, _ := ctx.FormFile("image")
	// Get filename postfix
	split := strings.Split(imageFile.Filename, ".")
	uuid := utils.GenerateUUID()
	postFix := split[len(split)-1]
	imageName := "/static/images/" + uuid + "." + postFix
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	price := ctx.PostForm("price")
	fullDescription := ctx.PostForm("full")
	token := ctx.PostForm("token")
	log.Println("token")
	log.Println(env.GetToken())
	if token != env.GetToken() {
		ctx.String(403, "Invalid Token")
		return
	}
	// Save Image File
	err := ctx.SaveUploadedFile(imageFile, "./static/images/"+uuid+"."+postFix)
	if err != nil {
		ctx.Error(err)
		return
	}
	// Save Item
	err = db.GetDb().CreateItem(model.Item{
		Id:           0,
		ItemImage:    imageName,
		ItemName:     name,
		ItemPrice:    price,
		ItemDesc:     description,
		ItemFullDesc: fullDescription,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Redirect(302, "/")
	return
}

func EditItemApi(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}
	item, err := db.GetDb().GetItemById(uint(idInt))
	if err != nil {
		ctx.Error(err)
		return
	}
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	price := ctx.PostForm("price")
	fullDescription := ctx.PostForm("full")
	token := ctx.PostForm("token")
	item.ItemName = name
	item.ItemDesc = description
	item.ItemPrice = price
	item.ItemFullDesc = fullDescription
	if token != env.GetToken() {
		ctx.String(403, "Invalid Token")
		return
	}
	err = db.GetDb().UpdateItem(item)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, item)
}

func DeleteItemApi(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.Error(err)
		return
	}
	token := ctx.Param("token")
	if token != env.GetToken() {
		ctx.String(403, "Invalid Token")
		return
	}
	// Delete related image
	item, err := db.GetDb().GetItemById(uint(idInt))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = os.Remove("." + item.ItemImage)
	err = db.GetDb().DelItem(uint(idInt))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.String(200, "OK")
}
