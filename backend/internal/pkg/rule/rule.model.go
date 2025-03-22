package rule

type Rule interface {
	GetType() Type
	GetInfo() BaseRule
	SetId(id string)
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

func (r *BaseRule) SetId(id string) {
	r.ID = id
}
