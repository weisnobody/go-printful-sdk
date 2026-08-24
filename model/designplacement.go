package model

type DesignPlacement struct {
	Placement        string          `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	Technique        string          `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	PrintAreaWidth   float64         `json:"print_area_width" bson:"print_area_width" mapstructure:"print_area_width" db:"print_area_width" csv:"print_area_width"`
	PrintAreaHeight  float64         `json:"print_area_height" bson:"print_area_height" mapstructure:"print_area_height" db:"print_area_height" csv:"print_area_height"`
	Layers           []FileLayer     `json:"layers" bson:"layers" mapstructure:"layers" db:"layers" csv:"layers"`
	PlacementOptions []CatalogOption `json:"placement_options" bson:"placement_options" mapstructure:"placement_options" db:"placement_options" csv:"placement_options"`
}
