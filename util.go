package mocksqs

import (
	"fmt"
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

// CreateQueueURL will generate a valid SQS queue URL with the provided queue
// name.
func CreateQueueURL(name string) string {
	return fmt.Sprintf("https://sqs.us-east-1.amazonaws.com/281910179584/%s", name)
}
