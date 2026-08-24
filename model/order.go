package model

type Order struct {
	ID          int     `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	ExternalID  string  `json:"external_id" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	StoreID     int     `json:"store_id" bson:"store_id" mapstructure:"store_id" db:"store_id" csv:"store_id"`
	Shipping    string  `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	Status      string  `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	CreatedAt   string  `json:"created_at" bson:"created_at" mapstructure:"created_at" db:"created_at" csv:"created_at"`
	UpdatedAt   string  `json:"updated_at" bson:"updated_at" mapstructure:"updated_at" db:"updated_at" csv:"updated_at"`
	Recipient   Address `json:"recipient" bson:"recipient" mapstructure:"recipient" db:"recipient" csv:"recipient"`
	Costs       `json:"costs" bson:"costs" mapstructure:"costs" db:"costs" csv:"costs"`
	RetailCosts `json:"retail_costs" bson:"retail_costs" mapstructure:"retail_costs" db:"retail_costs" csv:"retail_costs"`
	OrderItems  `json:"order_items" bson:"order_items" mapstructure:"order_items" db:"order_items" csv:"order_items"`
}
