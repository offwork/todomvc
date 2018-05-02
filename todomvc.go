package main

import (
	"todomvc/store"
	"todomvc/store/model"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type myComponent struct {
	vecty.Core
	Items        []*model.Item `vecty:"prop"`
	newItemTitle string
}

func (mc *myComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("my-main-container")),
			vecty.Text("Hadi LAN!"),
		),
	)
}

func main() {
	vecty.SetTitle("GopherJS • TodoMVC")
	vecty.AddStylesheet("todomvc/styles/base.css")
	vecty.AddStylesheet("todomvc/styles/index.css")

	comp := &myComponent{}
	store.Listeners.Add(comp, func() {})

	vecty.RenderBody(comp)
}
