package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type myComponent struct {
	vecty.Core
}

func (mc *myComponent) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			vecty.Markup(vecty.Class("my-main-container")),
			vecty.Text("Welcome to my site"),
		),
	)
}

func main() {
	comp := &myComponent{}
	vecty.RenderBody(comp)
}
