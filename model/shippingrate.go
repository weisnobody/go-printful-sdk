package model

type ShippingRate struct {
	Shipping           string      `json:"shipping" bson:"shipping" mapstructure:"shipping" db:"shipping" csv:"shipping"`
	ShippingMethodName string      `json:"shipping_method_name" bson:"shipping_method_name" mapstructure:"shipping_method_name" db:"shipping_method_name" csv:"shipping_method_name"`
	Rate               string      `json:"rate" bson:"rate" mapstructure:"rate" db:"rate" csv:"rate"`
	Currency           string      `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	MinDeliveryDays    int         `json:"min_delivery_days" bson:"min_delivery_days" mapstructure:"min_delivery_days" db:"min_delivery_days" csv:"min_delivery_days"`
	MaxDeliveryDays    int         `json:"max_delivery_days" bson:"max_delivery_days" mapstructure:"max_delivery_days" db:"max_delivery_days" csv:"max_delivery_days"`
	MinDeliveryDate    string      `json:"min_delivery_date" bson:"min_delivery_date" mapstructure:"min_delivery_date" db:"min_delivery_date" csv:"min_delivery_date"`
	MaxDeliveryDate    string      `json:"max_delivery_date" bson:"max_delivery_date" mapstructure:"max_delivery_date" db:"max_delivery_date" csv:"max_delivery_date"`
	Shipments          []Shipment2 `json:"shipments" bson:"shipments" mapstructure:"shipments" db:"shipments" csv:"shipments"`
}

type Shipment2 struct {
	DepartureCountry    string          `json:"departure_country" bson:"departure_country" mapstructure:"departure_country" db:"departure_country" csv:"departure_country"`
	ShipmentItems       []ShipmentItem2 `json:"shipment_items" bson:"shipment_items" mapstructure:"shipment_items" db:"shipment_items" csv:"shipment_items"`
	CustomsFeesPossible bool            `json:"customs_fees_possible" bson:"customs_fees_possible" mapstructure:"customs_fees_possible" db:"customs_fees_possible" csv:"customs_fees_possible"`
}

type ShipmentItem2 struct {
	CatalogVariantID int `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	Quantity         int `json:"quantity" bson:"quantity" mapstructure:"quantity" db:"quantity" csv:"quantity"`
}
