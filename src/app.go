package main

import (
        "github.com/ocampor/go-api/src/services"
	"github.com/ocampor/go-api/src/config"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
)

func createClient(uri string) *gorm.DB {
	db, err := gorm.Open("postgres", uri)

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	uri := os.Getenv("POSTGRES_URI")
	db := createClient(uri)

	u := services.UnitResource{db}
	restful.DefaultContainer.Add(u.UnitService())

	restful.DefaultContainer.Add(services.HealthCheckService())

	spec := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: config.SpecSwagger,
	}

	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(spec))

	//TODO: Replace hard-coded path to swagger dist
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("/Users/ricardo/Workspace/swagger-ui/dist"))))

	log.Printf("start listening on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
