package services

import (
	"../models"
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type UnitResource struct {
	Db *gorm.DB
}

func (u UnitResource) UnitService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/units").Produces(restful.MIME_JSON)

	ws.Route(
		ws.GET("/{unit-id}").To(u.findUnit).
			Doc("get unit details").
			Param(
				ws.PathParameter(
					"unit-id",
					"identifier of the property").
					DataType("integer")).
			Writes(models.Unit{}).
			Returns(200, "OK", models.Unit{}).
			Returns(404, "Not Found", nil))

	return ws
}

func (u UnitResource) findUnit(request *restful.Request, response *restful.Response) {
	id, _ := strconv.Atoi(request.PathParameter("unit-id"))

	var unit = models.Unit{}
	u.Db.First(&unit, id)

	if unit.PropertyId == 0 {
		response.WriteErrorString(http.StatusNotFound, "Unit could not be found")
	} else {
		response.WriteEntity(unit)
	}
}
