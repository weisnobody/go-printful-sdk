package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type VariantAvailabilityResponse struct {
	Data  []model.VariantAvailability `json:"data" bson:"data" mapstructure:"data" db:"data" csv:"data"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}
