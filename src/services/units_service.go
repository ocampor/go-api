package services

import (
	"../models"
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
)

type UnitResource struct {
	Db *gorm.DB
}

func (u UnitResource) UnitService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/units").Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/{property_id}").To(u.findUnit).
			Doc("get unit details").
			Param(
				ws.PathParameter("property_id", "identifier of the property").DataType("integer")).
			Writes(models.Unit{}).
			Returns(http.StatusOK, http.StatusText(http.StatusOK), models.Unit{}).
			Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil))

	ws.Route(
		ws.GET("/validate/{property_id}").To(u.validateDetails).
			Doc("compare details of a property with the official government records").
			Param(ws.PathParameter("property_id", "identifier of the property").DataType("integer")).
			Param(ws.QueryParameter("location_id", "property finder unit location identifier").DataType("integer")).
			Param(ws.QueryParameter("bedroom_id", "property finder bedrooms identifier").DataType("integer")).
			Param(ws.QueryParameter("unit_size", "plot area of the property").DataType("float32")).
			Param(ws.QueryParameter("unit_number", "unit number of the property").DataType("string")).
			Writes(models.DetailValidation{}).
			Returns(http.StatusOK, http.StatusText(http.StatusOK), models.DetailValidation{}).
			Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil))
	return ws
}

func (u UnitResource) findUnit(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("property_id")

	var unit = models.Unit{}
	u.Db.First(&unit, id)

	if unit.PropertyId == 0 {
		response.WriteErrorString(http.StatusNotFound, "Unit could not be found")
	} else {
		response.WriteEntity(unit)
	}
}

func (u UnitResource) validateDetails(request *restful.Request, response *restful.Response) {
	var detailValidation = models.DetailValidation{0.9, true, true, 0.84, 0.75}

	response.WriteEntity(detailValidation)
}
