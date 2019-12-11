package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// CreateQueue is not implemented. It will panic in all cases.
func (client *SQS) CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	panic("CreateQueue is not implemented")
}

// CreateQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) CreateQueueWithContext(aws.Context, *sqs.CreateQueueInput, ...request.Option) (*sqs.CreateQueueOutput, error) {
	panic("CreateQueueWithContext is not implemented")
}

// CreateQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) CreateQueueRequest(*sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput) {
	panic("CreateQueueRequest is not implemented")
}
