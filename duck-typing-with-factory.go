package main

import (
	"github.com/k0kubun/pp"
)

func main() {

	var layouts []Layout

	layouts = append(layouts, CreateLayout("text"))
	layouts = append(layouts, CreateLayout("multimedia"))

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

func CreateLayout(typeValue string) Layout {
	switch {
	case typeValue == "text":
		return &TextLayout{Type: "text"}
	case typeValue == "multimedia":
		return &MultiMediaLayout{Type: "multimedia"}
	}

	return &TextLayout{}
}
