package rule

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type Type struct {
	slug string
}

func (r *Type) Value() (driver.Value, error) {
	return strings.ToUpper(r.slug), nil
}

func (r *Type) Scan(src interface{}) error {
	var valueString string

	if _, ok := src.([]uint8); ok {
		valueString = string(src.([]uint8))
	} else {
		srcString := src.(string)
		valueString = srcString
	}

	r.slug = strings.ToUpper(valueString)

	return nil
}

func (r *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.slug)
}

func (r *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	r.slug = strings.ToUpper(s)
	return nil
}

var (
	Unknown = Type{""}
	Http    = Type{"HTTP"}
	Route   = Type{"ROUTE"}
)
