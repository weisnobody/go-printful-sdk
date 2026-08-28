package model

type PrintFileInfo struct {
	ProductID           int                `json:"product_id" bson:"product_id" mapstructure:"product_id" db:"product_id" csv:"product_id"`
	AvailablePlacements interface{}        `json:"available_placements" bson:"available_placements" mapstructure:"available_placements" db:"available_placements" csv:"available_placements"`
	Printfiles          []Printfile        `json:"printfiles" bson:"printfiles" mapstructure:"printfiles" db:"printfiles" csv:"printfiles"`
	VariantPrintfiles   []VariantPrintfile `json:"variant_printfiles" bson:"variant_printfiles" mapstructure:"variant_printfiles" db:"variant_printfiles" csv:"variant_printfiles"`
	OptionGroups        []string           `json:"option_groups" bson:"option_groups" mapstructure:"option_groups" db:"option_groups" csv:"option_groups"`
	Options             []string           `json:"options" bson:"options" mapstructure:"options" db:"options" csv:"options"`
}

type Printfile struct {
	PrintfileID int    `json:"printfile_id" bson:"printfile_id" mapstructure:"printfile_id" db:"printfile_id" csv:"printfile_id"`
	Width       int    `json:"width" bson:"width" mapstructure:"width" db:"width" csv:"width"`
	Height      int    `json:"height" bson:"height" mapstructure:"height" db:"height" csv:"height"`
	DPI         int    `json:"dpi" bson:"dpi" mapstructure:"dpi" db:"dpi" csv:"dpi"`
	FillMode    string `json:"fill_mode" bson:"fill_mode" mapstructure:"fill_mode" db:"fill_mode" csv:"fill_mode"`
	CanRotate   bool   `json:"can_rotate" bson:"can_rotate" mapstructure:"can_rotate" db:"can_rotate" csv:"can_rotate"`
}

type VariantPrintfile struct {
	VariantID  int            `json:"variant_id" bson:"variant_id" mapstructure:"variant_id" db:"variant_id" csv:"variant_id"`
	Placements map[string]int `json:"placements" bson:"placements" mapstructure:"placements" db:"placements" csv:"placements"`
}
