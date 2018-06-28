package router

type Router struct {

	Path string
	Method string
	Platform string
	Controller string
	Action string

}

func (r *Router) Parser() *Router {
	p := parser{router:r,path:r.Path}
	p.parser()
	return r
}

func (r *Router) Dispatch() {
	d := dispatch{router:r}
	d.dispatch()
}

func (r *Router) get(rule string) {

	r.setRule(rule)

}

func (r *Router) setRule(rule string) {

}
