package responses

import model "github.com/baldurstod/go-printful-sdk/model"

type Printfile struct {
	Code   int                 `json:"code" bson:"code" mapstructure:"code" db:"code" csv:"code"`
	Result model.PrintFileInfo `json:"result" bson:"result" mapstructure:"result" db:"result" csv:"result"`
}
