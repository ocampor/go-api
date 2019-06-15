package main

import (
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"log"
	"net/http"
)

func main() {
	webservice := new(restful.WebService)
	webservice.Route(
		webservice.GET("/health_check").
			To(isOk).
			Produces(restful.MIME_JSON))

	restful.DefaultContainer.Add(webservice)

	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/apidocs.json",
	}

	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))

	http.Handle(
		"/docs/",
		http.StripPrefix(
			"/docs/",
			http.FileServer(
				//TODO: Replace hard-coded path to swagger dist
				http.Dir("/Users/ricardo/Workspace/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func isOk(request *restful.Request, response *restful.Response) {
	response.WriteAsJson(map[string]string{"message": "OK"})
}
