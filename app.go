// +build js,wasm

package engine

import "syscall/js"

var App *Element

func CreateApp()  {
	App = NewElement("div")
}


func StartApp() {
	js.Global().Get("document").Get("body").Call("appendChild", App.Node)
}


