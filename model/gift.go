package model

type Gift struct {
	Subject string `json:"subject" bson:"subject" mapstructure:"subject" db:"subject" csv:"subject"`
	Message string `json:"message" bson:"message" mapstructure:"message" db:"message" csv:"message"`
}
