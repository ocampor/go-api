package services

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"strconv"
)

type UnitResource struct {
	Units map[int]Unit
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
			Writes(Unit{}).
			Returns(200, "OK", Unit{}).
			Returns(404, "Not Found", nil))

	return ws
}

func (u UnitResource) findUnit(request *restful.Request, response *restful.Response) {
	id, _ := strconv.Atoi(request.PathParameter("unit-id"))
	unit := u.Units[id]

	if unit.PropertyId == 0 {
		response.WriteErrorString(http.StatusNotFound, "Unit could not be found.")
	} else {
		response.WriteEntity(unit)
	}
}

type Unit struct {
	PropertyId int     `json:"property_id" description:"government database property identifier"`
	LocationId int     `json:"location_id" description:"property finder unit location identifier"`
	BedroomId  int     `json:"bedroom_id" description:"property finder bedrooms identifier"`
	UnitSize   float32 `json:"unit_size" description:"plot area of the property"`
	UnitNumber string  `json:"unit_number" description:"unit number of the property"`
}
