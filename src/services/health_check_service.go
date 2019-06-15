package services

import "github.com/emicklei/go-restful"

func HealthCheckService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/health_check").Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").
		To(isOk).
		Doc("check api status"))

	return ws
}

func isOk(request *restful.Request, response *restful.Response) {
	response.WriteAsJson(map[string]string{"message": "OK"})
}
