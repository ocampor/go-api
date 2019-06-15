package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
)

func main() {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("/health_check").To(isOk))

	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

func isOk(req *restful.Request, res *restful.Response) {
	io.WriteString(res, "OK")
}
