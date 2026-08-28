package model

type SyncProduct struct {
	ID           int64  `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	ExternalID   string `json:"external_id" bson:"external_id" mapstructure:"external_id" db:"external_id" csv:"external_id"`
	Name         string `json:"name" bson:"name" mapstructure:"name" db:"name" csv:"name"`
	Variants     int    `json:"variants" bson:"variants" mapstructure:"variants" db:"variants" csv:"variants"`
	Synced       int    `json:"synced" bson:"synced" mapstructure:"synced" db:"synced" csv:"synced"`
	Thumbnail    string `json:"thumbnail" bson:"thumbnail" mapstructure:"thumbnail" db:"thumbnail" csv:"thumbnail"`
	ThumbnailURL string `json:"thumbnail_url" bson:"thumbnail_url" mapstructure:"thumbnail_url" db:"thumbnail_url" csv:"thumbnail_url"`
	IsIgnored    bool   `json:"is_ignored" bson:"is_ignored" mapstructure:"is_ignored" db:"is_ignored" csv:"is_ignored"`
}

type SyncProductInfo struct {
	SyncProduct  SyncProduct   `json:"sync_product" bson:"sync_product" mapstructure:"sync_product" db:"sync_product" csv:"sync_product"`
	SyncVariants []SyncVariant `json:"sync_variants" bson:"sync_variants" mapstructure:"sync_variants" db:"sync_variants" csv:"sync_variants"`
}
