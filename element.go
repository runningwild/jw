package jw

import (
	"fmt"
	"syscall/js"
)

type Element struct {
	js.Value
}

func (e *Element) AppendChild(child *Element) {
	e.Call("appendChild", child.Value)
}
func (e *Element) RemoveChild(child *Element) {
	e.Call("removeChild", child.Value)
}

// Children returns the data returned by the element's e.children field, but unlike in javascript
// where this would be a live HTMLCollection, it is a non-live slice of those elements.
func (e *Element) Children() []*Element {
	var v []*Element
	children := e.Get("children")
	for i := 0; i < children.Get("length").Int(); i++ {
		v = append(v, &Element{children.Index(i)})
	}
	return v
}

func (e *Element) AddEventListener(event EventType, f func(this js.Value, args []js.Value) interface{}) {
	e.Call("addEventListener", string(event), js.FuncOf(f))
}
func (e *Element) Set(f FieldName, val interface{}) {
	e.Value.Set(string(f), val)
}
func (e *Element) GetElementsByClassName(className string) []*Element {
	v := e.Call("getElementsByClassName", className)
	fmt.Printf("Got back %v %d\n", v.Get("length"), v.Get("length").Int())
	var es []*Element
	for i := 0; i < v.Get("length").Int(); i++ {
		es = append(es, &Element{v.Call("item", i)})
	}
	return es
}
func (e *Element) GetClassList() *ClassList {
	return &ClassList{e.Get("classList")}
}
