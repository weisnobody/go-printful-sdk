package model

type VariantImages struct {
	CatalogVariantId  int     `json:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	Color             string  `json:"color" mapstructure:"color" db:"color" csv:"color"`
	PrimaryHexColor   string  `json:"primary_hex_color" mapstructure:"primary_hex_color" db:"primary_hex_color" csv:"primary_hex_color"`
	SecondaryHexColor string  `json:"secondary_hex_color" mapstructure:"secondary_hex_color" db:"secondary_hex_color" csv:"secondary_hex_color"`
	Images            []Image `json:"images" mapstructure:"images" db:"images" csv:"images"`
}

type Image struct {
	Placement       string `json:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	ImageUrl        string `json:"image_url" mapstructure:"image_url" db:"image_url" csv:"image_url"`
	BackgroundColor string `json:"background_color" mapstructure:"background_color" db:"background_color" csv:"background_color"`
	BackgroundImage string `json:"background_image" mapstructure:"background_image" db:"background_image" csv:"background_image"`
	MockupStyleId   int    `json:"mockup_style_id" mapstructure:"mockup_style_id" db:"mockup_style_id" csv:"mockup_style_id"`
}
