package model

type Item struct {
	ID             int            `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	ExternalID     string         `json:"external_id,omitempty" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	Quantity       int            `json:"quantity" bson:"quantity" mapstructure:"quantity" db:"quantity" csv:"quantity"`
	RetailPrice    string         `json:"retail_price" bson:"retail_price" mapstructure:"retail_price" db:"retail_price" csv:"retail_price"`
	Name           string         `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Placements     PlacementsList `json:"placements" bson:"placements" mapstructure:"placements" db:"placements" csv:"placements"`
	ProductOptions `json:"product_options,omitempty" bson:"product_options" mapstructure:"product_options" db:"product_options" csv:"product_options"`
}

type PlacementsList = []Placement

type ItemReadonly = Item
