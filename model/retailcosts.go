package model

type RetailCosts struct {
	CalculationStatus `json:"calculation_status" bson:"calculation_status" mapstructure:"calculation_status" db:"calculation_status" csv:"calculation_status"`
	Currency          string `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	Subtotal          string `json:"subtotal" bson:"subtotal" mapstructure:"subtotal" db:"subtotal" csv:"subtotal"`
	Discount          string `json:"discount" bson:"discount" mapstructure:"discount" db:"discount" csv:"discount"`
	Shipping          string `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	Vat               string `json:"vat" bson:"vat" mapstructure:"vat" db:"vat" csv:"vat"`
	Tax               string `json:"tax" bson:"tax" mapstructure:"tax" db:"tax" csv:"tax"`
	Total             string `json:"total" bson:"total" mapstructure:"total" db:"total" csv:"total"`
}
