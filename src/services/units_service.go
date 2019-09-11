package services

import (
	"../models"
	"../utils"
	"github.com/emicklei/go-restful"
	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

const (
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
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
			Produces(jsonapi.MediaType).
			Returns(http.StatusOK, http.StatusText(http.StatusOK), models.Unit{}).
			Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil))

	ws.Route(
		ws.GET("/validate/{property_id}").To(u.validateDetails).
			Doc("compare details of a property with the official government records").
			Param(ws.PathParameter("property_id", "identifier of the property").DataType("integer")).
			Param(ws.QueryParameter("location_id", "property finder unit location identifier").DataType("integer")).
			Param(ws.QueryParameter("bedroom_id", "property finder bedrooms identifier").DataType("integer")).
			Param(ws.QueryParameter("unit_size", "plot area of the property in square feet").DataType("float64")).
			Param(ws.QueryParameter("unit_number", "unit number of the property").DataType("string")).
			Writes(models.DetailValidation{}).
			Produces(jsonapi.MediaType).
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
		return
	}

	response.ResponseWriter.Header().Set(headerContentType, jsonapi.MediaType)
	response.ResponseWriter.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(response.ResponseWriter, &unit); err != nil {
		http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

func (u UnitResource) validateDetails(request *restful.Request, response *restful.Response) {
	id, err := strconv.Atoi(request.PathParameter("property_id"))

	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, "property_id must be an integer")
		return
	}

	locationId, _ := strconv.Atoi(request.QueryParameter("location_id"))
	bedroomId, _ := strconv.Atoi(request.QueryParameter("bedroom_id"))
	unitSize, _ := strconv.ParseFloat(request.QueryParameter("unit_size"), 64)
	unitNumber := request.QueryParameter("unit_number")

	var unit = models.Unit{}
	u.Db.First(&unit, id)

	if unit.PropertyId == 0 {
		response.WriteErrorString(http.StatusNotFound, "Unit could not be found")
		return
	}

	var unitSizeSqft = unitSize*10.7639

	//TODO: Extract this from the handler and move it to detail Validation
	locationSimilarity := utils.IntegerSimilarity(&locationId, unit.LocationId)
	bedroomSimilarity := utils.IntegerSimilarity(&bedroomId, unit.BedroomId)
	unitSizeSimilarity := utils.FloatSimilarity(unit.UnitSize, &unitSizeSqft)
	unitNumberSimilarity := utils.StringSimilarity(unit.UnitNumber, &unitNumber)
	overallSimilarity := utils.OverallSimilarity(
		locationSimilarity,
		bedroomSimilarity,
		unitSizeSimilarity,
		unitNumberSimilarity)

	detailValidation := models.DetailValidation{
		id,
		overallSimilarity,
		locationSimilarity,
		bedroomSimilarity,
		unitSizeSimilarity,
		unitNumberSimilarity,
	}

	response.ResponseWriter.Header().Set(headerContentType, jsonapi.MediaType)
	response.ResponseWriter.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(response.ResponseWriter, &detailValidation); err != nil {
		http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}
