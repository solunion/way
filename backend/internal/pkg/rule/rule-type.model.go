package rule

import (
	"database/sql/driver"
	"fmt"
)

type Type struct {
	slug string
}

func (r *Type) String() string {
	return r.slug
}

func (r *Type) Value() (driver.Value, error) {
	return r.slug, nil
}

func (r *Type) Scan(src interface{}) error {
	var valueString string

	if _, ok := src.([]uint8); ok {
		valueString = string(src.([]uint8))
	} else {
		srcString := src.(string)
		valueString = srcString
	}

	if value, err := TypeFromString(valueString); err != nil {
		return err
	} else {
		*r = value
	}

	return nil
}

var (
	Unknown = Type{""}
	Http    = Type{"HTTP"}
	Route   = Type{"ROUTE"}
)

func TypeFromString(s string) (Type, error) {
	switch s {
	case "HTTP":
		return Http, nil
	case "ROUTE":
		return Route, nil
	default:
		return Unknown, fmt.Errorf("unknown rule type: %s", s)
	}
}
