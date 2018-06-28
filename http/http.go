package http

import (
	"net/http"
	"github.com/liugene/webgo/router"
)

type FastHttp struct {}

func (*FastHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	route := router.Router{Path:r.URL.String()}
	route.Parser().Dispatch()
	req := Request{
		Path:r.URL.String(),
		Method:r.Method,
		Host:r.Host,
	}
	req.start()
}

func (ft *FastHttp) Listen(port string) {
	mux := &FastHttp{}
	http.ListenAndServe(":"+port, mux)
}

