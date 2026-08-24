package model

type Variant struct {
	ID               int            `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Name             string         `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	CatalogProductID int            `json:"catalog_product_id" bson:"catalog_product_id" mapstructure:"catalog_product_id" db:"catalog_product_id" csv:"catalog_product_id"`
	Color            string         `json:"color" bson:"color" mapstructure:"color" db:"color" csv:"color"`
	ColorCode        string         `json:"color_code" bson:"color_code" mapstructure:"color_code" db:"color_code" csv:"color_code"`
	ColorCode2       string         `json:"color_code2" bson:"color_code2" mapstructure:"color_code2" db:"color_code_2" csv:"color_code_2"`
	Image            string         `json:"image" bson:"image" mapstructure:"image" db:"image" csv:"image"`
	Size             string         `json:"size" bson:"size" mapstructure:"size" db:"size" csv:"size"`
	Availability     []Availability `json:"availability" bson:"availability" mapstructure:"availability" db:"availability" csv:"availability"`
}

type Availability struct {
	Region string `json:"region" bson:"region" mapstructure:"region" db:"region" csv:"region"`
	Status string `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
}
