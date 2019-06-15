package services

import (
	"../models"
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"math"
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
			Param(ws.QueryParameter("unit_size", "plot area of the property").DataType("float64")).
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

func floatSimilarity(a float64, b float64) float64 {
	if a < b {
		return math.Abs(a / b)
	} else {
		return math.Abs(b / a)
	}
}

func boolToFloat64(a bool) float64 {
	if a {
		return 1.0
	} else {
		return 0.0
	}
}

func stringSimilarity(a string, b string) float64 {
	return levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
}

func (u UnitResource) validateDetails(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("property_id")
	locationId, _ := strconv.Atoi(request.QueryParameter("location_id"))
	bedroomId, _ := strconv.Atoi(request.QueryParameter("bedroom_id"))
	unitSize, _ := strconv.ParseFloat(request.QueryParameter("unit_size"), 64)
	unitNumber := request.QueryParameter("unit_number")

	var unit = models.Unit{}
	u.Db.First(&unit, id)

	if unit.PropertyId == 0 {
		response.WriteErrorString(http.StatusNotFound, "Unit could not be found")
	}

	locationMatches := locationId == unit.LocationId
	bedroomMatches := bedroomId == unit.BedroomId

	unitSizeSimilarity := floatSimilarity(unit.UnitSize, unitSize)
	unitNumberSimilarity := stringSimilarity(unit.UnitNumber, unitNumber)

	overallSimilarity := (boolToFloat64(locationMatches) +
		boolToFloat64(bedroomMatches) +
		unitSizeSimilarity +
		unitNumberSimilarity) / 4

	var detailValidation = models.DetailValidation{
		overallSimilarity,
		locationMatches,
		bedroomMatches,
		unitSizeSimilarity,
		unitNumberSimilarity,
	}

	response.WriteEntity(detailValidation)
}
