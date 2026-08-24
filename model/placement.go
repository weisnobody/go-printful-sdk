package model

type Placement struct {
	Placement         string  `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	Technique         string  `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	PrintAreaType     string  `json:"print_area_type" bson:"print_area_type" mapstructure:"print_area_type" db:"print_area_type" csv:"print_area_type"`
	Layers            []Layer `json:"layers" bson:"layers" mapstructure:"layers" db:"layers" csv:"layers"`
	PlacementOptions  `json:"placement_options" bson:"placement_options" mapstructure:"placement_options" db:"placement_options" csv:"placement_options"`
	Status            string `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	StatusExplanation string `json:"status_explanation" bson:"status_explanation" mapstructure:"status_explanation" db:"status_explanation" csv:"status_explanation"`
}

type PlacementOptions []PlacementOption

type PlacementOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values" db:"values" csv:"values"`
}

func NewPlacement() Placement {
	return Placement{
		PrintAreaType: "simple",
	}
}
