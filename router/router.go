package router

import (
	"net/http"
)

type router struct {

}

func (router *router) get(rule string) {

	router.setRule(rule)

}

func (router *router) setRule(rule string) {

	mux := http.NewServeMux()
	mux.Handle(rule)

}
