package router

import (
	"application/http/controller"
)

type dispatch struct {
	router *Router
}

func (d *dispatch) dispatch() {
	pack := controller.Home{}
	pack.Index()
}