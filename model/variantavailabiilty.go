package model

type VariantAvailability struct {
	CatalogVariantID int                     `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	Techniques       []AvailabilityTechnique `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
}

type AvailabilityTechnique struct {
	Technique      string                      `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	SellingRegions []AvailabilitySellingRegion `json:"selling_regions" bson:"selling_regions" mapstructure:"selling_regions" db:"selling_regions" csv:"selling_regions"`
}

type AvailabilitySellingRegion struct {
	Name                        string                        `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Availability                string                        `json:"availability" bson:"availability" mapstructure:"availability" db:"availability" csv:"availability"`
	PlacementOptionAvailability []AvailabilityPlacementOption `json:"placement_option_availability" bson:"placement_option_availability" mapstructure:"placement_option_availability" db:"placement_option_availability" csv:"placement_option_availability"`
}

type AvailabilityPlacementOption struct {
	OptionName string `json:"option_name" bson:"option_name" mapstructure:"option_name" db:"option_name" csv:"option_name"`
	Available  bool   `json:"available" bson:"available" mapstructure:"available" db:"available" csv:"available"`
}
