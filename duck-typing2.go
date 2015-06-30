package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	var layouts []Layout

	layouts = append(layouts, &TextLayout{Type: "text"})
	layouts = append(layouts, &MultiMediaLayout{Type: "multimedia"})

	for _, layout := range layouts {
		pp.Println(layout.GetType())
	}
}

type Layout interface {
	GetType() string
}

type TextLayout struct {
	Type string
}

func (tl *TextLayout) GetType() string {
	return tl.Type
}

type MultiMediaLayout struct {
	Type string
}

func (ml *MultiMediaLayout) GetType() string {
	return ml.Type
}
