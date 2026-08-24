package model

type Country struct {
	Name   string  `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Code   string  `json:"code" bson:"code" mapstructure:"code" db:"code" csv:"code"`
	Region string  `json:"region" bson:"region" mapstructure:"region" db:"region" csv:"region"`
	States []State `json:"states" bson:"states" mapstructure:"states" db:"states" csv:"states"`
}

type State struct {
	Name string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Code string `json:"code" bson:"code" mapstructure:"code" db:"code" csv:"code"`
}
