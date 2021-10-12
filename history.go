package jw

import "syscall/js"

type History struct {
	js.Value
}

var hist History

func init() {
	hist.Value = js.Global().Get("history")
}
func GetHistory() *History {
	return &hist
}

func (h *History) Back() {
	h.Value.Call("back")
}
func (h *History) Forward() {
	h.Value.Call("forward")
}
func (h *History) Go(n int) {
	h.Value.Call("go", n)
}
func (h *History) Len() int {
	return h.Value.Get("length").Int()
}
func (h *History) PushState(obj interface{}, url string) {
	h.Value.Call("pushState", obj, "", url)
}
func (h *History) ReplaceState(obj interface{}, url string) {
	h.Value.Call("replaceState", obj, "", url)
}
