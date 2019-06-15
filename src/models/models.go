package models

type Unit struct {
	PropertyId int     `json:"property_id" description:"government database property identifier" gorm:"primary_key"`
	LocationId int     `json:"location_id" description:"property finder unit location identifier"`
	BedroomId  int     `json:"bedroom_id" description:"property finder bedrooms identifier"`
	UnitSize   float32 `json:"unit_size" description:"plot area of the property"`
	UnitNumber string  `json:"unit_number" description:"unit number of the property"`
}

func (Unit) TableName() string {
	return "unit"
}
