// +build js,wasm

package engine

import "syscall/js"

type Element struct {
	Node js.Value
	child []*Element
	draw func()
}

func NewElement(tag string) *Element {
	e := &Element{}
	e.Node = js.Global().Get("document").Call("createElement", tag)
	e.child = make([]*Element , 0)
	return e
}

func (e *Element) SetInnerHtml(innerhtml string) {
	e.Node.Set("innerHTML", innerhtml)
}

func (e *Element) AddChild(element *Element) {
	e.Node.Call("appendChild", element.Node)
	//e.child = append(e.child , element)
}

func (e *Element) RemoveSingleChild(element *Element) {
	//e.child = nil
	e.Node.Call("removeChild",element.Node)

}

func (e *Element) RemoveChild() {
	//e.child = nil
	for {
		n := e.Node.Get("firstChild")
		if n == js.Null() {
			break
		}
		e.Node.Call("removeChild",n)
	}

}

func (e *Element) SetClass(class string) {
	e.Node.Set("className", class)
}

func (e *Element) SetId(id string) {
	e.Node.Set("id", id)
}

func (e *Element) SetCallBack(cbType  string , cb js.Callback) {
	e.Node.Call("addEventListener", cbType, cb)
}

func (e *Element) Set(key string , value string) {
	e.Node.Set(key,value)
}


