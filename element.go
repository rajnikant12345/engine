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

func (e *Element) SetInnerHtml(innerhtml string) *Element {
	e.Node.Set("innerHTML", innerhtml)
	return e
}

func (e *Element) AddChild(element *Element) *Element {
	e.Node.Call("appendChild", element.Node)
	return e
	//e.child = append(e.child , element)
}

func (e *Element) RemoveSingleChild(element *Element) *Element {
	//e.child = nil
	e.Node.Call("removeChild",element.Node)
	return e

}

func (e *Element) RemoveChild() *Element {
	//e.child = nil
	for {
		n := e.Node.Get("firstChild")
		if n == js.Null() {
			break
		}
		e.Node.Call("removeChild",n)
	}
	return e
}

func (e *Element) SetClass(class string) *Element {
	e.Node.Set("className", class)
	return e
}

func (e *Element) AppendClass(class string) *Element {
	c := e.Node.Get("className").String()
	c += " "+class
	e.Node.Set("className", c)
	return e
}

func (e *Element) SetStyle(key string, value string) *Element {
	e.Node.Get("style").Set(key,value)
	return e
}

func (e *Element) SetId(id string) *Element {
	e.Node.Set("id", id)
	return e
}

func (e *Element) SetCallBack(cbType  string , cb js.Callback) *Element {
	e.Node.Call("addEventListener", cbType, cb)
	return e
}

func (e *Element) Set(key string , value string) *Element {
	e.Node.Set(key,value)
	return e
}


func (e *Element) Nest(el ...*Element ) *Element {
	for _,i := range el {
		e.AddChild(i)
	}
	return e
}

