package webgo

import (
	"github.com/liugene/webgo/http"
)

type Web struct {
	name string
}

func (w *Web) Run() {
	h := new(http.FastHttp)
	h.Listen("8888")
}

func (w *Web) Router() {
}

func (w *Web) Event() {
}
