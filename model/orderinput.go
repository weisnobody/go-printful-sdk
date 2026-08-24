package model

type OrderInput struct {
	ExternalID    string        `json:"external_id" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	Shipping      string        `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	Recipient     Address       `json:"recipient" bson:"recipient" mapstructure:"recipient" db:"recipient" csv:"recipient"`
	OrderItems    []CatalogItem `json:"order_items" bson:"order_items" mapstructure:"order_items" db:"order_items" csv:"order_items"`
	Customization Customization `json:"customization" bson:"customization" mapstructure:"customization" db:"customization" csv:"customization"`
	RetailCosts   RetailCosts2  `json:"retail_costs" bson:"retail_costs" mapstructure:"retail_costs" db:"retail_costs" csv:"retail_costs"`
}
