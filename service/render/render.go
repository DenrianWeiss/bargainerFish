package render

import (
	"github.com/gomarkdown/markdown"
)

func RenderRichText(payload []byte) []byte {
	return markdown.ToHTML(payload, nil, nil)
}
