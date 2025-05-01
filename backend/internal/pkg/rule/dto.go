package rule

type Common struct {
	ID          string  `mapstructure:"id" json:"id"`
	Type        string  `mapstructure:"type" json:"type"`
	Name        string  `mapstructure:"name" json:"name"`
	Description *string `mapstructure:"description" json:"description"`
}

type CreateRequest struct {
	Common
	ID    string                 `json:"-"`
	Value map[string]interface{} `json:"value"`
}

type UpdateRequest struct {
	Common
	Value map[string]interface{} `json:"value"`
}

type Response struct {
	Common Common `mapstructure:", squash"`
}
