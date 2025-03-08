package rule

type Rule interface {
	GetType() Type
	GetInfo() BaseRule
}

type BaseRule struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	Type         Type    `json:"-"`
	TypeInString string  `json:"type"`
}

func (r *BaseRule) GetType() Type {
	return r.Type
}

func (r *BaseRule) GetInfo() BaseRule {
	return *r
}
