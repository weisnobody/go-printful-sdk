package responses

type Error4XXResponse struct {
	Code   int    `json:"code" bson:"code" mapstructure:"code" db:"code" csv:"code"`
	Result string `json:"result" bson:"result" mapstructure:"result" db:"result" csv:"result"`
	Error  `json:"error" bson:"error" mapstructure:"error" db:"error" csv:"error"`
}

type Error struct {
	Reason  string `json:"reason" bson:"reason" mapstructure:"reason" db:"reason" csv:"reason"`
	Message string `json:"message" bson:"message" mapstructure:"message" db:"message" csv:"message"`
}

type Error5XXResponse struct {
	Type     string `json:"type" bson:"type" mapstructure:"type" db:"type" csv:"type"`
	Status   int    `json:"status" bson:"status" mapstructure:"status" db:"status" csv:"status"`
	Title    string `json:"title" bson:"title" mapstructure:"title" db:"title" csv:"title"`
	Details  string `json:"details" bson:"details" mapstructure:"details" db:"details" csv:"details"`
	Instance string `json:"instance" bson:"instance" mapstructure:"instance" db:"instance" csv:"instance"`
}
