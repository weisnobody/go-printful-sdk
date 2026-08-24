package model

type ShippingRatesAddress struct {
	Address1    string `json:"address1" bson:"address1" mapstructure:"address1" db:"address_1" csv:"address_1"`
	Address2    string `json:"address2" bson:"address2" mapstructure:"address2" db:"address_2" csv:"address_2"`
	City        string `json:"city" bson:"city" mapstructure:"city" db:"city" csv:"city"`
	StateCode   string `json:"state_code" bson:"state_code" mapstructure:"state_code" db:"state_code" csv:"state_code"`
	CountryCode string `json:"country_code" bson:"country_code" mapstructure:"country_code" db:"country_code" csv:"country_code"`
	ZIP         string `json:"zip" bson:"zip" mapstructure:"zip" db:"zip" csv:"zip"`
}
