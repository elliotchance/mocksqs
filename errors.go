package mocksqs

import "fmt"

func errorMissingField(field string) error {
	return fmt.Errorf("InvalidParameter: 1 validation error(s) found.\n- missing required field, %s.\n", field)
}

func errorMissingParameter(field string) error {
	return errorWithRequestID(fmt.Sprintf("MissingParameter: The request must contain the parameter %s.\n", field))
}

func errorInvalidParameterValue(message string) error {
	return errorWithRequestID(fmt.Sprintf("InvalidParameterValue: %s", message))
}

func errorInternal() error {
	return errorWithRequestID(fmt.Sprintf("InternalError: We encountered an internal error. Please try again."))
}

func errorNonExistentQueue() error {
	return errorWithRequestID("AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.")
}

func errorWithRequestID(message string) error {
	return fmt.Errorf("%s\n\tstatus code: 400, request id: %s", message, newRequestID())
}
