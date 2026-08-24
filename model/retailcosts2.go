package model

type RetailCosts2 struct {
	Currency string `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	Discount string `json:"discount" bson:"discount" mapstructure:"discount" db:"discount" csv:"discount"`
	Shipping string `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	Tax      string `json:"tax" bson:"tax" mapstructure:"tax" db:"tax" csv:"tax"`
}
