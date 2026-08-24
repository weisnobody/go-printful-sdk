package model

// Note: Dimensions should be int as declared in openapi.json, but the actual API send floats
// BackgroundColor should be int but is actually a string
type MockupTemplates struct {
	CatalogVariantIDs   []int   `json:"catalog_variant_ids" bson:"catalog_variant_ids" mapstructure:"catalog_variant_i_ds" db:"catalog_variant_i_ds" csv:"catalog_variant_i_ds"`
	Placement           string  `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	Technique           string  `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	ImageURL            string  `json:"image_url" bson:"image_url" mapstructure:"image_url" db:"image_url" csv:"image_url"`
	BackgroundURL       string  `json:"background_url" bson:"background_url" mapstructure:"background_url" db:"background_url" csv:"background_url"`
	BackgroundColor     string  `json:"background_color" bson:"background_color" mapstructure:"background_color" db:"background_color" csv:"background_color"`
	PrintfileID         int     `json:"printfile_id" bson:"printfile_id" mapstructure:"printfile_id" db:"printfile_id" csv:"printfile_id"`
	TemplateWidth       float64 `json:"template_width" bson:"template_width" mapstructure:"template_width" db:"template_width" csv:"template_width"`
	TemplateHeight      float64 `json:"template_height" bson:"template_height" mapstructure:"template_height" db:"template_height" csv:"template_height"`
	PrintAreaWidth      float64 `json:"print_area_width" bson:"print_area_width" mapstructure:"print_area_width" db:"print_area_width" csv:"print_area_width"`
	PrintAreaHeight     float64 `json:"print_area_height" bson:"print_area_height" mapstructure:"print_area_height" db:"print_area_height" csv:"print_area_height"`
	PrintAreaTop        float64 `json:"print_area_top" bson:"print_area_top" mapstructure:"print_area_top" db:"print_area_top" csv:"print_area_top"`
	PrintAreaLeft       float64 `json:"print_area_left" bson:"print_area_left" mapstructure:"print_area_left" db:"print_area_left" csv:"print_area_left"`
	TemplatePositioning string  `json:"template_positioning" bson:"template_positioning" mapstructure:"template_positioning" db:"template_positioning" csv:"template_positioning"`
	Orientation         string  `json:"orientation" bson:"orientation" mapstructure:"orientation" db:"orientation" csv:"orientation"`
}
