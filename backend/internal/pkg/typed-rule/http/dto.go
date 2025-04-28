package http

import (
	"encoding/json"
	"github.com/solunion/way/backend/internal/pkg/typed-rule/generic"
)

type CreateRequest struct {
	Response
	ID string `json:"-"`
}

func (r *CreateRequest) Value() HttpRuleValue {
	return HttpRuleValue{
		Method: r.Method,
		Path:   r.Path,
	}
}

type HttpRuleValueResponse struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Response struct {
	generic.Response
	HttpRuleValueResponse
}

func (r *Response) FromModel(rule *HttpRule) error {
	if marshal, err := json.Marshal(rule); err != nil {
		return err
	} else {
		if err := json.Unmarshal(marshal, r); err != nil {
			return err
		}

		if marshalValue, err := json.Marshal(rule.Value); err != nil {
			return err
		} else {
			if err := json.Unmarshal(marshalValue, &r.HttpRuleValueResponse); err != nil {
				return err
			}
		}
	}
	return nil
}
