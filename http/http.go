package http

import "net/http"

type FastHttp struct {
}

func ServeHttp(w http.ResponseWriter, r http.Request) {

}

func (ft *FastHttp) Listen(port string) {
	http.ListenAndServe(":"+port, nil)
}

