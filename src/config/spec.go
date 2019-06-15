package config

import "github.com/go-openapi/spec"

func SpecSwagger(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Dubai Land Department Service",
			Description: "Resource to retrieve information from land department databases",
			Contact: &spec.ContactInfo{
				Name:  "Ricardo Ocampo",
				Email: "ricardo.ocampo@propertyfinder.ae",
			},
			Version: "1.0.0-alpha",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "units",
		Description: "Land Department units details"}}}
}
