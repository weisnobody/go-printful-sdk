package model

type Customization struct {
	Gift        `json:"gift" bson:"gift" mapstructure:"gift" db:"gift" csv:"gift"`
	PackingSlip `json:"packing_slip" bson:"packing_slip" mapstructure:"packing_slip" db:"packing_slip" csv:"packing_slip"`
}
