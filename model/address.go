package model

type Address struct {
	Name        string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Company     string `json:"company" bson:"company" mapstructure:"company" db:"company" csv:"company"`
	Address1    string `json:"address1" bson:"address1" mapstructure:"address1" db:"address_1" csv:"address_1"`
	Address2    string `json:"address2" bson:"address2" mapstructure:"address2" db:"address_2" csv:"address_2"`
	City        string `json:"city" bson:"city" mapstructure:"city" db:"city" csv:"city"`
	StateCode   string `json:"state_code" bson:"state_code" mapstructure:"state_code" db:"state_code" csv:"state_code"`
	StateName   string `json:"state_name" bson:"state_name" mapstructure:"state_name" db:"state_name" csv:"state_name"`
	CountryCode string `json:"country_code" bson:"country_code" mapstructure:"country_code" db:"country_code" csv:"country_code"`
	CountryName string `json:"country_name" bson:"country_name" mapstructure:"country_name" db:"country_name" csv:"country_name"`
	ZIP         string `json:"zip" bson:"zip" mapstructure:"zip" db:"zip" csv:"zip"`
	Phone       string `json:"phone" bson:"phone" mapstructure:"phone" db:"phone" csv:"phone"`
	Email       string `json:"email" bson:"email" mapstructure:"email" db:"email" csv:"email"`
	TaxNumber   string `json:"tax_number" bson:"tax_number" mapstructure:"tax_number" db:"tax_number" csv:"tax_number"`
}
