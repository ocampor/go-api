package models

type Unit struct {
	PropertyId int     `jsonapi:"primary,unit" description:"government database property identifier" gorm:"primary_key"`
	LocationId int     `jsonapi:"attr,location_id" description:"property finder unit location identifier"`
	BedroomId  int     `jsonapi:"attr,bedroom_id" description:"property finder bedrooms identifier"`
	UnitSize   float64 `jsonapi:"attr,unit_size" description:"plot area of the property"`
	UnitNumber string  `jsonapi:"attr,unit_number" description:"unit number of the property"`
}

type DetailValidation struct {
	PropertyId           int     `jsonapi:"primary,detail_validation" description:"government database property identifier"`
	OverallSimilarity    float64 `jsonapi:"attr,overall_similarity" description:"overall similarity between goverment details and listing details"`
	LocationMatches      float64 `jsonapi:"attr,location_similarity" description:"the listing is under the same location tree as the government property"`
	BedroomMatches       float64 `jsonapi:"attr,bedroom_similarity" description:"the listing bedrooms and the government property matches"`
	UnitSizeSimilarity   float64 `jsonapi:"attr,unit_size_similarity" description:"similarity between government property unit size and listing"`
	UnitNumberSimilarity float64 `jsonapi:"attr,unit_number_similarity" description:"similarity the listing unit number and the government property"`
}

func (Unit) TableName() string {
	return "unit"
}
