package responses

type Paging struct {
	Total  uint `json:"total" bson:"total" mapstructure:"total" db:"total" csv:"total"`
	Offset uint `json:"offset" bson:"offset" mapstructure:"offset" db:"offset" csv:"offset"`
	Limit  uint `json:"limit" bson:"limit" mapstructure:"limit" db:"limit" csv:"limit"`
}
