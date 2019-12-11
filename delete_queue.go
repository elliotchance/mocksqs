package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// DeleteQueue is not implemented. It will panic in all cases.
func (client *SQS) DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	panic("DeleteQueue is not implemented")
}

// DeleteQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) DeleteQueueWithContext(aws.Context, *sqs.DeleteQueueInput, ...request.Option) (*sqs.DeleteQueueOutput, error) {
	panic("DeleteQueueWithContext is not implemented")
}

// DeleteQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) DeleteQueueRequest(*sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput) {
	panic("DeleteQueueRequest is not implemented")
}
