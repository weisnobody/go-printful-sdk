package model

type MockupStyles struct {
	Placement       string        `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	DisplayName     string        `json:"display_name" bson:"display_name" mapstructure:"display_name" db:"display_name" csv:"display_name"`
	Technique       string        `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	PrintAreaWidth  float64       `json:"print_area_width" bson:"print_area_width" mapstructure:"print_area_width" db:"print_area_width" csv:"print_area_width"`
	PrintAreaHeight float64       `json:"print_area_height" bson:"print_area_height" mapstructure:"print_area_height" db:"print_area_height" csv:"print_area_height"`
	PrintAreaType   string        `json:"print_area_type" bson:"print_area_type" mapstructure:"print_area_type" db:"print_area_type" csv:"print_area_type"`
	Dpi             int           `json:"dpi" bson:"dpi" mapstructure:"dpi" db:"dpi" csv:"dpi"`
	MockupStyles    []MockupStyle `json:"mockup_styles" bson:"mockup_styles" mapstructure:"mockup_styles" db:"mockup_styles" csv:"mockup_styles"`
}

type MockupStyle struct {
	Id                   int    `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	CategoryName         string `json:"category_name" bson:"category_name" mapstructure:"category_name" db:"category_name" csv:"category_name"`
	ViewName             string `json:"view_name" bson:"view_name" mapstructure:"view_name" db:"view_name" csv:"view_name"`
	RestrictedToVariants []int  `json:"restricted_to_variants" bson:"restricted_to_variants" mapstructure:"restricted_to_variants" db:"restricted_to_variants" csv:"restricted_to_variants"`
}
