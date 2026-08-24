package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type ProductImagesResponse struct {
	Data   []model.VariantImages `json:"data" mapstructure:"data" db:"data" csv:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}
