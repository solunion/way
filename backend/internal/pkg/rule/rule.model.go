package rule

import (
	"encoding/json"
	"fmt"
)

type HttpRule struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Method      string  `json:"method"`
	Path        string  `json:"path"`
}

func (r *HttpRule) Type() Type {
	return Http
}

func (r *HttpRule) Value() []byte {
	var value = map[string]interface{}{
		"method": r.Method,
		"path":   r.Path,
	}

	result, err := json.Marshal(value)

	if err != nil {
		fmt.Println(err)
	}

	return result
}

func (r *HttpRule) FromValue(value []byte) error {
	return json.Unmarshal(value, r)
}
