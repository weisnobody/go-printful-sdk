package model

type CatalogItemType string

const (
	TypeOrderItem    CatalogItemType = "order_item"
	TypeBrandingItem CatalogItemType = "branding_item"
)

type CatalogItemSummary struct {
	ID               int             `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Type             CatalogItemType `json:"Type" bson:"Type" mapstructure:"type" db:"type" csv:"type"`
	CatalogVariantID int             `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	ExternalID       string          `json:"external_id" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	Quantity         int             `json:"quantity" bson:"quantity" mapstructure:"quantity" db:"quantity" csv:"quantity"`
	Name             string          `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Price            string          `json:"price" bson:"price" mapstructure:"price" db:"price" csv:"price"`
	RetailPrice      string          `json:"retail_price" bson:"retail_price" mapstructure:"retail_price" db:"retail_price" csv:"retail_price"`
	Currency         string          `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	RetailCurrency   string          `json:"retail_currency" bson:"retail_currency" mapstructure:"retail_currency" db:"retail_currency" csv:"retail_currency"`
}

func (o *CatalogItemSummary) isOrderItem() {}
