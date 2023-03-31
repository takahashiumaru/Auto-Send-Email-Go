package helper

import (
	"strings"
)

func ErrorRequestMessage(err error) string {
	var message = "Error : "
	errors := []DuplicateError{
		// CITY
		{
			contains: "CityCreateRequest.ID",
			message:  "Field validation for id required, max 3 character",
		},
		{
			contains: "CityCreateRequest.Name",
			message:  "Field validation for id required, max 100 character",
		},
		{
			contains: "CityUpdateRequest.ID",
			message:  "Field validation for id required, max 3 character",
		},
		{
			contains: "CityUpdateRequest.Name",
			message:  "Field validation for id required, max 100 character",
		},
	}

	for _, duplicateError := range errors {
		if strings.Contains(err.Error(), duplicateError.contains) {
			message += duplicateError.message + " | "
		}
	}
	if message != "Error : " {
		message += "XXX"
		message = strings.ReplaceAll(message, "| XXX", "")
	} else {
		message += err.Error()
	}

	return message
}
