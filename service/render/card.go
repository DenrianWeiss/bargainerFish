package render

import (
	"bytes"
	"github.com/DenrianWeiss/barginerFish/model"
	"strconv"
	"text/template"
)

type RenderInput struct {
	Id       string `json:"id"`
	Image    string `json:"image"`
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Content  string `json:"content"`
}

var cardTemplate *template.Template

func init() {
	templateStr := `<div class="mdui-card card" style="break-inside: avoid; border-radius: 14px; overflow: auto;">
  <div class="mdui-card-media">
    <img src="{{.Image}}"/>
  </div>
  <div class="mdui-card-primary">
    <div class="mdui-card-primary-title">{{.Title}}</div>
    <div class="mdui-card-primary-subtitle">{{.SubTitle}}</div>
  </div>
  <div class="mdui-card-content">{{.Content}}</div>
  <div class="mdui-card-actions">
    <button class="mdui-btn mdui-ripple" style="border-radius: 14px;" onclick="window.open('items/{{.Id}}')">Details</button>
    </button>
  </div>
</div>`
	parse, err := template.New("card").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	cardTemplate = parse
}

func ConvertItemToRenderCard(i model.Item) RenderInput {
	r := RenderInput{
		Id:       strconv.Itoa(int(i.Id)),
		Image:    i.ItemImage,
		Title:    i.ItemName,
		SubTitle: i.ItemPrice,
		Content:  string(RenderRichText([]byte(i.ItemDesc))),
	}
	return r
}

func RenderCard(input RenderInput) string {
	buf := bytes.NewBufferString("")
	err := cardTemplate.Execute(buf, input)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
