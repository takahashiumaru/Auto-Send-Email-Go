package helper

import (
	"errors"
)

func OperatorQuery(symbol string) (string, error) {
	var operator string
	switch symbol {
	case "eq":
		operator = "="
	case "like":
		operator = "like"
	case "lt":
		operator = "<"
	case "lte":
		operator = "<="
	case "gt":
		operator = ">"
	case "gte":
		operator = ">="
	case "ne":
		operator = "<>"
	case "bw":
		operator = "BETWEEN"
	case "in":
		operator = "IN"
	default:
		return operator, errors.New("operator symbol paramater is not valid")
	}
	return operator, nil
}
