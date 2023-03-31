package helper

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplyFilter(tx *gorm.DB, filters *map[string]string) error {
	for key, value := range *filters {
		keySplitted := strings.Split(key, ".")
		column := keySplitted[0]
		operator, err := OperatorQuery(keySplitted[1])
		if err != nil {
			return err
		}
		query := fmt.Sprintf("%s %s ?", column, operator)
		switch operator {
		case "like":
			value = "%" + value + "%"
			tx = tx.Where(query, value)
		case "BETWEEN":
			columnSplitted := strings.Split(column, "|")
			tx = tx.Where("? "+operator+" "+columnSplitted[0]+" AND "+columnSplitted[1], value)
		case "IN":
			valueSplitted := strings.Split(value, ",")
			tx = tx.Where(query, valueSplitted)
		default:
			tx = tx.Where(query, value)
		}
	}
	return nil
}
