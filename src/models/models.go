package models

type Unit struct {
	PropertyId int     `json:"property_id" description:"government database property identifier" gorm:"primary_key"`
	LocationId int     `json:"location_id" description:"property finder unit location identifier"`
	BedroomId  int     `json:"bedroom_id" description:"property finder bedrooms identifier"`
	UnitSize   float64 `json:"unit_size" description:"plot area of the property"`
	UnitNumber string  `json:"unit_number" description:"unit number of the property"`
}

type DetailValidation struct {
	OverallSimilarity    float64 `json:"overall_similarity" description:"overall similarity between goverment details and listing details"`
	LocationMatches      bool    `json:"location_matches" description:"the listing is under the same location tree as the government property"`
	BedroomMatches       bool    `json:"bedroom_matches" description:"the listing bedrooms and the government property matches"`
	UnitSizeSimilarity   float64 `json:"unit_size_similarity" description:"similarity between government property unit size and listing"`
	UnitNumberSimilarity float64 `json:"unit_number_similarity" description:"similarity the listing unit number and the government property"`
}

func (Unit) TableName() string {
	return "unit"
}
