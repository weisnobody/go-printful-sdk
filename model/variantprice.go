package model

type VariantPrice struct {
	Currency string              `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	Product  ProductPriceInfo    `json:"product" bson:"product" mapstructure:"product" db:"product" csv:"product"`
	Variant  VariantsPriceData `json:"variant" bson:"variant" mapstructure:"variant" db:"variant" csv:"variant"`
}

type ProductPriceInfo struct {
	ID         int                    `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Placements []AdditionalPlacements `json:"placements" bson:"placements" mapstructure:"placements" db:"placements" csv:"placements"`
}

type AdditionalPlacements struct {
	ID               string             `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Title            string             `json:"title" bson:"title" mapstructure:"title" db:"title" csv:"title"`
	Type             string             `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	TechniqueKey     string             `json:"technique_key" bson:"technique_key" mapstructure:"technique_key" db:"technique_key" csv:"technique_key"`
	Price            string             `json:"price" bson:"price" mapstructure:"price" db:"price" csv:"price"`
	DiscountedPrice  string             `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price" db:"discounted_price" csv:"discounted_price"`
	PlacementOptions []FileOptionPrices `json:"placement_options" bson:"placement_options" mapstructure:"placement_options" db:"placement_options" csv:"placement_options"`
	Layers           []Layers           `json:"layers" bson:"layers" mapstructure:"layers" db:"layers" csv:"layers"`
}

type Layers struct {
	Type            string              `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	AdditionalPrice string              `json:"additional_price" bson:"additional_price" mapstructure:"additional_price" db:"additional_price" csv:"additional_price"`
	Options         []LayerOptionPrices `json:"layer_options" bson:"layer_options" mapstructure:"layer_options" db:"options" csv:"options"`
}

type LayerOptionPrices struct {
	Name        string            `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Type        string            `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Values      []any             `json:"values" bson:"values" mapstructure:"values" db:"values" csv:"values"`
	Description string            `json:"description" bson:"description" mapstructure:"description" db:"description" csv:"description"`
	Price       map[string]string `json:"price" bson:"price" mapstructure:"price" db:"price" csv:"price"`
}

type VariantsPriceData struct {
	ID         int                  `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Techniques []TechniquePriceInfo `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
}

type TechniquePriceInfo struct {
	Price           string `json:"price" bson:"price" mapstructure:"price" db:"price" csv:"price"`
	DiscountedPrice string `json:"discounted_price" bson:"discounted_price" mapstructure:"discounted_price" db:"discounted_price" csv:"discounted_price"`
	TechniqueKey    string `json:"technique_key" bson:"technique_key" mapstructure:"technique_key" db:"technique_key" csv:"technique_key"`
	DisplayName     string `json:"technique_display_name" bson:"technique_display_name" mapstructure:"technique_display_name" db:"display_name" csv:"display_name"`
}

type FileOptionPrices struct {
	Name        string            `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Type        string            `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Values      []any             `json:"values" bson:"values" mapstructure:"values" db:"values" csv:"values"`
	Description string            `json:"description" bson:"description" mapstructure:"description" db:"description" csv:"description"`
	Price       map[string]string `json:"price" bson:"price" mapstructure:"price" db:"price" csv:"price"`
}
