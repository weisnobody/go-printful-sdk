package model

type FileLayer struct {
	Type         string          `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	LayerOptions []CatalogOption `json:"layer_options" bson:"layer_options" mapstructure:"layer_options" db:"layer_options" csv:"layer_options"`
}
