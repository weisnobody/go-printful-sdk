package model

type ProductInfo struct {
	Product `bson:"inline" mapstructure:",squash" db:"product" csv:"product"`
}

type Product struct {
	ID                int                `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	MainCategoryID    int                `json:"main_category_id" bson:"main_category_id" mapstructure:"main_category_id" db:"main_category_id" csv:"main_category_id"`
	Categories        []int              `json:"categories" bson:"categories" mapstructure:"categories" db:"categories" csv:"categories"`
	Type              string             `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Name              string             `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Brand             string             `json:"brand" bson:"brand" mapstructure:"brand" db:"brand" csv:"brand"`
	Model             string             `json:"model" bson:"model" mapstructure:"model" db:"model" csv:"model"`
	Image             string             `json:"image" bson:"image" mapstructure:"image" db:"image" csv:"image"`
	ImageWomen        string             `json:"image_women" bson:"image_women" mapstructure:"image_women" db:"image_women" csv:"image_women"`
	VariantCount      int                `json:"variant_count" bson:"variant_count" mapstructure:"variant_count" db:"variant_count" csv:"variant_count"`
	CatalogVariantIDs []int              `json:"catalog_variant_ids" bson:"catalog_variant_ids" mapstructure:"catalog_variant_i_ds" db:"catalog_variant_i_ds" csv:"catalog_variant_i_ds"`
	IsDiscontinued    bool               `json:"is_discontinued" bson:"is_discontinued" mapstructure:"is_discontinued" db:"is_discontinued" csv:"is_discontinued"`
	Description       string             `json:"description" bson:"description" mapstructure:"description" db:"description" csv:"description"`
	Sizes             []string           `json:"sizes" bson:"sizes" mapstructure:"sizes" db:"sizes" csv:"sizes"`
	Colors            []Color            `json:"colors" bson:"colors" mapstructure:"colors" db:"colors" csv:"colors"`
	Techniques        []Technique        `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
	Placements        []ProductPlacement `json:"placements" bson:"placements" mapstructure:"placements" db:"placements" csv:"placements"`
	ProductOptions    []CatalogOption    `json:"product_options" bson:"product_options" mapstructure:"product_options" db:"product_options" csv:"product_options"`
}

type ProductPlacement struct {
	DesignPlacement       `bson:"inline" mapstructure:"design_placement" db:"design_placement" csv:"design_placement"`
	ConflictingPlacements []string `json:"conflicting_placements" bson:"conflicting_placements" mapstructure:"conflicting_placements" db:"conflicting_placements" csv:"conflicting_placements"`
}

type CatalogOption struct {
	Name       string   `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Techniques []string `json:"techniques" bson:"techniques" mapstructure:"techniques" db:"techniques" csv:"techniques"`
	Type       string   `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Values     any      `json:"values" bson:"values" mapstructure:"values" db:"values" csv:"values"`
}
