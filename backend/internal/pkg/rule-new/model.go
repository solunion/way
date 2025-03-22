package rule_new

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Rule interface {
	SetID(id uuid.UUID)
	GetID() uuid.UUID
	GetType() Type
	GetName() string
	GetDescription() *string
	GetValue() (json.RawMessage, error)
}

type BasicRule struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Type        Type            `json:"type"`
	Value       json.RawMessage `json:"-"`
}

func (r *BasicRule) SetID(id uuid.UUID) {
	r.ID = id
}

func (r *BasicRule) GetID() uuid.UUID {
	return r.ID
}

func (r *BasicRule) GetType() Type {
	return r.Type
}

func (r *BasicRule) GetName() string {
	return r.Name
}

func (r *BasicRule) GetDescription() *string {
	return r.Description
}

func (r *BasicRule) GetValue() (json.RawMessage, error) {
	return r.Value, nil
}
