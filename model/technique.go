package model

type Technique struct {
	Key         string `json:"key" bson:"key" mapstructure:"key" db:"key" csv:"key"`
	DisplayName string `json:"display_name" bson:"display_name" mapstructure:"display_name" db:"display_name" csv:"display_name"`
	IsDefault   bool   `json:"is_default" bson:"is_default" mapstructure:"is_default" db:"is_default" csv:"is_default"`
}
