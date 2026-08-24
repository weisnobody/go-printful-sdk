package model

type PackingSlip struct {
	Email         string `json:"email" bson:"email" mapstructure:"email" db:"email" csv:"email"`
	Phone         string `json:"phone" bson:"phone" mapstructure:"phone" db:"phone" csv:"phone"`
	Message       string `json:"message" bson:"message" mapstructure:"message" db:"message" csv:"message"`
	LogoURL       string `json:"logo_url" bson:"logo_url" mapstructure:"logo_url" db:"logo_url" csv:"logo_url"`
	StoreName     string `json:"store_name" bson:"store_name" mapstructure:"store_name" db:"store_name" csv:"store_name"`
	CustomOrderID string `json:"custom_order_id" bson:"custom_order_id" mapstructure:"custom_order_id" db:"custom_order_id" csv:"custom_order_id"`
}
