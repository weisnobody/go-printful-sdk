package responses

type MockupTasksResponse struct {
	Data   []MockupGeneratorTask `json:"data" bson:"data" mapstructure:"data" db:"data" csv:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}

type MockupGeneratorTask struct {
	ID                    int                    `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Status                string                 `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	CatalogVariantMockups []CatalogVariantMockup `json:"catalog_variant_mockups" bson:"catalog_variant_mockups" mapstructure:"catalog_variant_mockups" db:"catalog_variant_mockups" csv:"catalog_variant_mockups"`
	FailureReasons        []FailureReasons       `json:"failure_reasons" bson:"failure_reasons" mapstructure:"failure_reasons" db:"failure_reasons" csv:"failure_reasons"`
}

type CatalogVariantMockup struct {
	CatalogVariantID int                     `json:"catalog_variant_id" bson:"catalog_variant_id" mapstructure:"catalog_variant_id" db:"catalog_variant_id" csv:"catalog_variant_id"`
	Mockups          []MockupGeneratorMockup `json:"mockups" bson:"mockups" mapstructure:"mockups" db:"mockups" csv:"mockups"`
}

type MockupGeneratorMockup struct {
	Placement   string `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	DisplayName string `json:"display_name" bson:"display_name" mapstructure:"display_name" db:"display_name" csv:"display_name"`
	Technique   string `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	StyleID     int    `json:"style_id" bson:"style_id" mapstructure:"style_id" db:"style_id" csv:"style_id"`
	MockupURL   string `json:"mockup_url" bson:"mockup_url" mapstructure:"mockup_url" db:"mockup_url" csv:"mockup_url"`
}

type FailureReasons struct {
	Type        string        `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Detail      string        `json:"detail" bson:"detail" mapstructure:"detail" db:"detail" csv:"detail"`
	Source      FailureSource `json:"source" bson:"source" mapstructure:"source" db:"source" csv:"source"`
	ValidValues []string      `json:"valid_values" bson:"valid_values" mapstructure:"valid_values" db:"valid_values" csv:"valid_values"`
}

type FailureSource struct {
	Header string `json:"header" bson:"header" mapstructure:"header" db:"header" csv:"header"`
}
