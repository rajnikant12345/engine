// +build js,wasm

package engine

import "syscall/js"

var App *Element


//CreateApp  is called by the main function to start your wasm app.
func CreateApp()  {
	App = NewElement("div")
}


//StartApp is clled by main app to start the applicatio and
//e.g. engine.CreateApp()
//	engine.App.AddChild( app.CreateApp() )
//	engine.StartApp()
func StartApp() {
	js.Global().Get("document").Get("body").Call("appendChild", App.Node)
}


