package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// DeleteQueue is fully supported.
func (client *SQS) DeleteQueue(input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	if input.QueueUrl == nil {
		return &sqs.DeleteQueueOutput{},
			errorMissingField("DeleteQueueInput.QueueUrl")
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		client.queues.Delete(*input.QueueUrl)

		return &sqs.DeleteQueueOutput{}, nil
	}

	return &sqs.DeleteQueueOutput{},
		errorWithRequestID("AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.")
}

// DeleteQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) DeleteQueueWithContext(aws.Context, *sqs.DeleteQueueInput, ...request.Option) (*sqs.DeleteQueueOutput, error) {
	panic("DeleteQueueWithContext is not implemented")
}

// DeleteQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) DeleteQueueRequest(*sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput) {
	panic("DeleteQueueRequest is not implemented")
}
