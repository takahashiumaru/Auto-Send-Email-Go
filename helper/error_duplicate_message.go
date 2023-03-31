package helper

import (
	"strings"
)

type DuplicateError struct {
	contains string
	message  string
}

func ErrorDuplicateMessage(err error) string {
	errors := []DuplicateError{
		{
			contains: "cities.PRIMARY",
			message:  "city already exists",
		},
		{
			contains: "outlets.npwp",
			message:  "npwp already exists",
		},
		{
			contains: "marketing_structures.idx_marketing_structure_code_period",
			message:  "code already exists in the same period",
		},
	}

	for _, duplicateError := range errors {
		if strings.Contains(err.Error(), duplicateError.contains) {
			return duplicateError.message
		}
	}

	return "record already exists"
}
