package model

type CalculationStatus string

const (
	Done        CalculationStatus = "done"
	Calculating CalculationStatus = "calculating"
	Failed      CalculationStatus = "failed"
)

type Costs struct {
	CalculationStatus `json:"calculation_status" bson:"calculation_status" mapstructure:"calculation_status" db:"calculation_status" csv:"calculation_status"`
	Currency          string `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	Subtotal          string `json:"subtotal" bson:"subtotal" mapstructure:"subtotal" db:"subtotal" csv:"subtotal"`
	Discount          string `json:"discount" bson:"discount" mapstructure:"discount" db:"discount" csv:"discount"`
	Shipping          string `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	Digitization      string `json:"digitization" bson:"digitization" mapstructure:"digitization" db:"digitization" csv:"digitization"`
	AdditionalFee     string `json:"additional_fee" bson:"additional_fee" mapstructure:"additional_fee" db:"additional_fee" csv:"additional_fee"`
	FulfillmentFee    string `json:"fulfillment_fee" bson:"fulfillment_fee" mapstructure:"fulfillment_fee" db:"fulfillment_fee" csv:"fulfillment_fee"`
	RetailDeliveryFee string `json:"retail_delivery_fee" bson:"retail_delivery_fee" mapstructure:"retail_delivery_fee" db:"retail_delivery_fee" csv:"retail_delivery_fee"`
	Vat               string `json:"vat" bson:"vat" mapstructure:"vat" db:"vat" csv:"vat"`
	Tax               string `json:"tax" bson:"tax" mapstructure:"tax" db:"tax" csv:"tax"`
	Total             string `json:"total" bson:"total" mapstructure:"total" db:"total" csv:"total"`
}
