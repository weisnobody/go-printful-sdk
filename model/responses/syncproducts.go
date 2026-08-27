package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type SyncProductsResponse struct {
	Code   int                 `json:"code" mapstructure:"code" db:"code" csv:"code"`
	Result []model.SyncProduct `json:"result" bson:"result" mapstructure:"result" db:"result" csv:"result"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}

type SyncProductResponse struct {
	Code   int                   `json:"code" mapstructure:"code" db:"code" csv:"code"`
	Result model.SyncProductInfo `json:"result" bson:"result" mapstructure:"result" db:"result" csv:"result"`
	Paging `json:"paging" bson:"paging" mapstructure:"paging" db:"paging" csv:"paging"`
}
