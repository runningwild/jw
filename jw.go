package jw

import (
	"syscall/js"
)

type Document struct {
	js.Value
}

var doc Document

func init() {
	doc.Value = js.Global().Get("document")
}
func GetDocument() *Document {
	return &doc
}

func (d *Document) GetElementByID(id string) *Element {
	e := d.Call("getElementById", id)
	if e == nil {
		return nil
	}
	return &Element{e}
}
func (d *Document) CreateElement(t TagName) *Element {
	return &Element{Value: d.Call("createElement", string(t))}
}

// func (d *document) getElementsByClassName(className string) []*Element {
// 	d.Call("getElementsByClassName")
// }

type TagName string

const (
	TagNameDiv TagName = "div"
)

type FieldName string

const (
	FieldNameInnerHTML FieldName = "innerHTML"
	FieldNameSrc       FieldName = "src"
)

type EventType string

const (
	EventTypeClick      EventType = "click"
	EventTypeMouseEnter EventType = "mouseenter"
	EventTypeMouseLeave EventType = "mouseleave"
	EventTypeMouseMove  EventType = "mousemove"
)

type ClassList struct {
	js.Value
}

func (cl *ClassList) Add(class string)     { cl.Call("add", class) }
func (cl *ClassList) Remove(class string)  { cl.Call("remove", class) }
func (cl *ClassList) Replace(class string) { cl.Call("replace", class) }
func (cl *ClassList) Toggle(class string)  { cl.Call("toggle", class) }
