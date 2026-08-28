package model

// for v2 of the API

type MockupTask struct {
	Format        string              `json:"format" bson:"format" mapstructure:"format" db:"format" csv:"format"`
	MockupWidthPX int                 `json:"mockup_width_px" bson:"mockup_width_px" mapstructure:"mockup_width_px" db:"mockup_width_px" csv:"mockup_width_px"`
	Products      []MockupTaskProduct `json:"products" bson:"products" mapstructure:"products" db:"products" csv:"products"`
}

// not using the Placement type from placement.go because the Layer type it references doesn't match the layer type
type MockupTaskProduct struct {
	Source            string                    `json:"source" bson:"source" mapstructure:"source" db:"source" csv:"source"`
	MockupStyleIDs    []int                     `json:"mockup_style_ids" bson:"mockup_style_ids" mapstructure:"mockup_style_i_ds" db:"mockup_style_i_ds" csv:"mockup_style_i_ds"`
	CatalogProductID  int                       `json:"catalog_product_id" bson:"catalog_product_id" mapstructure:"catalog_product_id" db:"catalog_product_id" csv:"catalog_product_id"`
	CatalogVariantIDs []int                     `json:"catalog_variant_ids" bson:"catalog_variant_ids" mapstructure:"catalog_variant_i_ds" db:"catalog_variant_i_ds" csv:"catalog_variant_i_ds"`
	Orientation       string                    `json:"orientation" bson:"orientation" mapstructure:"orientation" db:"orientation" csv:"orientation"`
	Placements        []MockupTaskPlacement     `json:"placements" bson:"placements" mapstructure:"placements" db:"placements" csv:"placements"`
	ProductOptions    []MockupTaskProductOption `json:"product_options" bson:"product_options" mapstructure:"product_options" db:"product_options" csv:"product_options"`
}

type MockupTaskPlacement struct {
	Placement        string                             `json:"placement" bson:"placement" mapstructure:"placement" db:"placement" csv:"placement"`
	Technique        string                             `json:"technique" bson:"technique" mapstructure:"technique" db:"technique" csv:"technique"`
	PrintAreaType    string                             `json:"print_area_type" bson:"print_area_type" mapstructure:"print_area_type" db:"print_area_type" csv:"print_area_type"`
	Layers           []MockupTaskPlacementLayer         `json:"layers" bson:"layers" mapstructure:"layers" db:"layers" csv:"layers"`
	PlacementOptions []MockupTaskProductPlacementOption `json:"placement_options" bson:"placement_options" mapstructure:"placement_options" db:"placement_options" csv:"placement_options"`
}

type MockupTaskPlacementLayer struct {
	Type         string                           `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	URL          string                           `json:"url" bson:"url" mapstructure:"url" db:"url" csv:"url"`
	LayerOptions []MockupTaskPlacementLayerOption `json:"layer_options" bson:"layer_options" mapstructure:"layer_options" db:"layer_options" csv:"layer_options"`
	Position     LayerPosition                    `json:"position" bson:"position" mapstructure:"position" db:"position" csv:"position"`
}

type MockupTaskPlacementLayerOption struct {
	Name  string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Value bool   `json:"value" bson:"value" mapstructure:"value" db:"value" csv:"value"`
}

/* // LayerPosition exists
type MockupTaskProductPlacementLayerPosition struct {
	Width	int	`json:"width" bson:"width"`
	Height	int	`json:"height" bson:"height"`
	Top	int	`json:"top" bson:"top"`
	Left	int	`json:"left" bson:"left"`
}
*/

type MockupTaskProductPlacementOption struct {
	Name  string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Value bool   `json:"value" bson:"value" mapstructure:"value" db:"value" csv:"value"`
}
type MockupTaskProductOption struct {
	Name  string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Value bool   `json:"value" bson:"value" mapstructure:"value" db:"value" csv:"value"`
}
