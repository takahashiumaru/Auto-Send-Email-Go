package helper

import (
	"reflect"
)

func GetJSONFields(i interface{}) []string {
	fields := []string{}

	val := reflect.ValueOf(i)
	for i := 0; i < val.Type().NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Tag.Get("json"))
	}
	return fields
}

func MessageDataFoundOrNot(data interface{}) string {
	s := reflect.ValueOf(data)

	if s.Kind() == reflect.Slice {
		if s.Len() > 0 {
			return "Record found"
		} else {
			return "Record not found"
		}
	}
	if s.Kind() == reflect.Struct {
		return "Record found"
	}
	if data == nil {
		return "Record not found"
	}
	panic("MessageData FoundOrNor() given parameter must be slice or struct")

}
