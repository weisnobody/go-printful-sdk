package model

type Color struct {
	Name  string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Value string `json:"value" bson:"value" mapstructure:"value" db:"value" csv:"value"`
}
