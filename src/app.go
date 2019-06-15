package main

import (
	"./config"
	"./services"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"log"
	"net/http"
)

func main() {
	restful.DefaultContainer.Add(services.HealthCheckService())

	spec := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: config.SpecSwagger,
	}

	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(spec))

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
