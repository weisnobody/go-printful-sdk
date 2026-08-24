package model

type CatalogItem struct {
	Source           string `json:"source" bson:"source" mapstructure:"source" db:"source" csv:"source"`
	CatalogVariantID int    `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	Item             `mapstructure:",squash" db:"item" csv:"item"`
}

func NewCatalogItem() CatalogItem {
	return CatalogItem{
		Source: "catalog",
	}
}

type CatalogItemReadonly struct {
	Source           string `json:"source" bson:"source" mapstructure:"source" db:"source" csv:"source"`
	CatalogVariantID int    `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	ItemReadonly     `mapstructure:",squash" db:"item_readonly" csv:"item_readonly"`
}
