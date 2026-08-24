package model

type OrderItems []OrderItem

type OrderItemItemSource string

const (
	SourceCatalog   OrderItemItemSource = "catalog"
	SourceWarehouse OrderItemItemSource = "warehouse"
)

type OrderItem struct {
	Source             string `json:"source" bson:"source" mapstructure:"source" db:"source" csv:"source"`
	CatalogItemSummary `mapstructure:"catalog_item_summary" db:"catalog_item_summary" csv:"catalog_item_summary"`
	//TODO: add WarehouseItemSummary
}
