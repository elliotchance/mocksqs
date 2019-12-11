package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// PurgeQueue is not implemented. It will panic in all cases.
func (client *SQS) PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	panic("PurgeQueue is not implemented")
}

// PurgeQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) PurgeQueueWithContext(aws.Context, *sqs.PurgeQueueInput, ...request.Option) (*sqs.PurgeQueueOutput, error) {
	panic("PurgeQueueWithContext is not implemented")
}

// PurgeQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) PurgeQueueRequest(*sqs.PurgeQueueInput) (*request.Request, *sqs.PurgeQueueOutput) {
	panic("PurgeQueueRequest is not implemented")
}
