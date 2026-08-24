package model

type Layer struct {
	Type           string `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Url            string `json:"url" bson:"url" mapstructure:"url" db:"url" csv:"url"`
	*LayerOptions  `json:"layer_options,omitempty" bson:"layer_options" mapstructure:"layer_options,omitempty"`
	*LayerPosition `json:"position,omitempty" bson:"position" mapstructure:"position,omitempty"`
}

type LayerOptions []LayerOption

type LayerOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values" db:"values" csv:"values"`
}

type LayerPosition struct {
	Width  float64 `json:"width" bson:"width" mapstructure:"width" db:"width" csv:"width"`
	Height float64 `json:"height" bson:"height" mapstructure:"height" db:"height" csv:"height"`
	Top    float64 `json:"top" bson:"top" mapstructure:"top" db:"top" csv:"top"`
	Left   float64 `json:"left" bson:"left" mapstructure:"left" db:"left" csv:"left"`
}
