package model

type ProductPrices struct {
	Currency string              `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	Product  ProductPriceInfo    `json:"product" bson:"product" mapstructure:"product" db:"product" csv:"product"`
	Variants []VariantsPriceData `json:"variants" bson:"variants" mapstructure:"variants" db:"variants" csv:"variants"`
}
