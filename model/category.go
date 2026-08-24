package model

type Category struct {
	ID       int    `json:"id" bson:"id" mapstructure:"id" db:"id" csv:"id"`
	ParentID int    `json:"parent_id" bson:"parent_id" mapstructure:"parent_id" db:"parent_id" csv:"parent_id"`
	ImageURL string `json:"image_url" bson:"image_url" mapstructure:"image_url" db:"image_url" csv:"image_url"`
	Title    string `json:"title" bson:"title" mapstructure:"title" db:"title" csv:"title"`
}
