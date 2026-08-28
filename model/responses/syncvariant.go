package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type SyncVariantResponse struct {
	Code   int                   `json:"code" mapstructure:"code" db:"code" csv:"code"`
	Result model.SyncVariantInfo `json:"result" bson:"result" mapstructure:"result" db:"result" csv:"result"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}
