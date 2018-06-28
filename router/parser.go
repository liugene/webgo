package router

import (
	"strings"
)

type parser struct {
	router *Router
	path string
}

func (p *parser) parser() {
	pa := strings.Split(p.path, "/")
	p.router.Platform = pa[1]
	p.router.Controller = pa[2]
	p.router.Action = pa[3]
}