package model

import (
	"encoding/json"
)

type SyncVariant struct {
	ID                        int64             `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	ExternalID                string            `json:"external_id" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	SyncProductID             int64             `json:"sync_product_id" bson:"sync_product_id" mapstructure:"sync_product_id" db:"sync_product_id" csv:"sync_product_id"`
	Name                      string            `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Synced                    bool              `json:"synced" bson:"synced" mapstructure:"synced" db:"synced" csv:"synced"`
	VariantID                 int               `json:"variant_id" bson:"variant_id" mapstructure:"variant_id" db:"variant_id" csv:"variant_id"`
	RetailPrice               string            `json:"retail_price" bson:"retail_price" mapstructure:"retail_price" db:"retail_price" csv:"retail_price"`
	Currency                  string            `json:"currency" bson:"currency" mapstructure:"currency" db:"currency" csv:"currency"`
	IsIgnored                 bool              `json:"is_ignored" bson:"is_ignored" mapstructure:"is_ignored" db:"is_ignored" csv:"is_ignored"`
	SKU                       string            `json:"sku" bson:"sku" mapstructure:"sku" db:"sku" csv:"sku"`
	Product                   ProductVariant    `json:"product" bson:"product" mapstructure:"product" db:"product" csv:"product"`
	Files                     []SyncVariantFile `json:"files" bson:"files" mapstructure:"files" db:"files" csv:"files"`
	Options                   []ItemOption      `json:"options" bson:"options" mapstructure:"options" db:"options" csv:"options"`
	MainCategoryID            int               `json:"main_category_id" bson:"main_category_id" mapstructure:"main_category_id" db:"main_category_id" csv:"main_category_id"`
	WarehouseProductVariantID int               `json:"warehouse_product_variant_id" bson:"warehouse_product_variant_id" mapstructure:"warehouse_product_variant_id" db:"warehouse_product_variant_id" csv:"warehouse_product_variant_id"`
	Size                      string            `json:"size" bson:"size" mapstructure:"size" db:"size" csv:"size"`
	Color                     string            `json:"color" bson:"color" mapstructure:"color" db:"color" csv:"color"`
	AvailabilityStatus        string            `json:"availability_status" bson:"availability_status" mapstructure:"availability_status" db:"availability_status" csv:"availability_status"`
}

type SyncVariantFile struct {
	Type            string       `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	ID              int          `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	URL             string       `json:"url" bson:"url" mapstructure:"url" db:"url" csv:"url"`
	Options         []FileOption `json:"options" bson:"options" mapstructure:"options" db:"options" csv:"options"`
	Hash            string       `json:"hash" bson:"hash" mapstructure:"hash" db:"hash" csv:"hash"`
	Filename        string       `json:"filename" bson:"filename" mapstructure:"filename" db:"filename" csv:"filename"`
	MimeType        string       `json:"mime_type" bson:"mime_type" mapstructure:"mime_type" db:"mime_type" csv:"mime_type"`
	Size            int          `json:"size" bson:"size" mapstructure:"size" db:"size" csv:"size"`
	Width           int          `json:"width" bson:"width" mapstructure:"width" db:"width" csv:"width"`
	Height          int          `json:"height" bson:"height" mapstructure:"height" db:"height" csv:"height"`
	DPI             int          `json:"dpi" bson:"dpi" mapstructure:"dpi" db:"dpi" csv:"dpi"`
	Status          string       `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	Created         int64        `json:"created" bson:"created" mapstructure:"created" db:"created" csv:"created"`
	ThumbnailURL    string       `json:"thumbnail_url" bson:"thumbnail_url" mapstructure:"thumbnail_url" db:"thumbnail_url" csv:"thumbnail_url"`
	PreviewURL      string       `json:"preview_url" bson:"preview_url" mapstructure:"preview_url" db:"preview_url" csv:"preview_url"`
	Visible         bool         `json:"visible" bson:"visible" mapstructure:"visible" db:"visible" csv:"visible"`
	IsTemporary     bool         `json:"is_temporary" bson:"is_temporary" mapstructure:"is_temporary" db:"is_temporary" csv:"is_temporary"`
	StitchCountTier string       `json:"stitch_count_tier" bson:"stitch_count_tier" mapstructure:"stitch_count_tier" db:"stitch_count_tier" csv:"stitch_count_tier"`
}

type ItemOption struct {
	ID       string      `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	RawValue interface{} `json:"value" bson:"value" mapstructure:"raw_value" db:"raw_value" csv:"raw_value"`
	Value    string      `json:"value_adj" bson:"value_adj" mapstructure:"value_adj" db:"value_adj" csv:"value_adj"`
}

type FileOption struct {
	ID    string `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	Value string `json:"value" bson:"value" mapstructure:"value" db:"value" csv:"value"`
}

type ProductVariant struct {
	VariantID int    `json:"variant_id" bson:"variant_id" mapstructure:"variant_id" db:"variant_id" csv:"variant_id"`
	ProductID int    `json:"product_id" bson:"product_id" mapstructure:"product_id" db:"product_id" csv:"product_id"`
	Image     string `json:"image" bson:"image" mapstructure:"image" db:"image" csv:"image"`
	Name      string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
}

type SyncVariantInfo struct {
	SyncProduct SyncProduct `json:"sync_product" bson:"sync_product" mapstructure:"sync_product" db:"sync_product" csv:"sync_product"`
	SyncVariant SyncVariant `json:"sync_variant" bson:"sync_variant" mapstructure:"sync_variant" db:"sync_variant" csv:"sync_variant"`
}

func (itemopt *ItemOption) UnmarshalJSON(data []byte) error {
	type localItem ItemOption
	var loc localItem
	//fmt.Println("UJ: ", mof.RawJSONValue)
	if err := json.Unmarshal(data, &loc); err != nil {
		return err
	}
	*itemopt = ItemOption(loc)
	switch itemopt.RawValue.(type) {
	case string:
		strval := itemopt.RawValue.(string)
		itemopt.Value = strval
		itemopt.RawValue = strval
	case []string:
		vals := itemopt.RawValue.([]interface{})
		if len(vals) > 0 {
			strval, _ := vals[0].(string)
			for _, v := range vals[1:] {
				strval += "," + v.(string)
			}
			itemopt.Value = strval
		}
	}
	return nil
}
