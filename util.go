package mocksqs

import (
	"github.com/google/uuid"
	"reflect"
)

func newRequestID() string {
	return uuid.New().String()
}

func checkRequiredFields(fields map[string]interface{}) error {
	for field, value := range fields {
		if reflect.ValueOf(value).IsNil() {
			return errorMissingField(field)
		}
	}

	return nil
}
