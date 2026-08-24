package model

type File struct {
	ID           int    `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	URL          string `json:"url" bson:"url" mapstructure:"url" db:"url" csv:"url"`
	Hash         string `json:"hash" bson:"hash" mapstructure:"hash" db:"hash" csv:"hash"`
	Filename     string `json:"filename" bson:"filename" mapstructure:"filename" db:"filename" csv:"filename"`
	MimeType     string `json:"mime_type" bson:"mime_type" mapstructure:"mime_type" db:"mime_type" csv:"mime_type"`
	Size         int    `json:"size" bson:"size" mapstructure:"size" db:"size" csv:"size"`
	Width        int    `json:"width" bson:"width" mapstructure:"width" db:"width" csv:"width"`
	Height       int    `json:"height" bson:"height" mapstructure:"height" db:"height" csv:"height"`
	Dpi          int    `json:"dpi" bson:"dpi" mapstructure:"dpi" db:"dpi" csv:"dpi"`
	Status       string `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	Created      string `json:"created" bson:"created" mapstructure:"created" db:"created" csv:"created"`
	ThumbnailURL string `json:"thumbnail_url" bson:"thumbnail_url" mapstructure:"thumbnail_url" db:"thumbnail_url" csv:"thumbnail_url"`
	PreviewURL   string `json:"preview_url" bson:"preview_url" mapstructure:"preview_url" db:"preview_url" csv:"preview_url"`
	Visible      bool   `json:"visible" bson:"visible" mapstructure:"visible" db:"visible" csv:"visible"`
	IsTemporary  bool   `json:"is_temporary" bson:"is_temporary" mapstructure:"is_temporary" db:"is_temporary" csv:"is_temporary"`
}
