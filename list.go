package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ListQueues is not implemented. It will panic in all cases.
func (client *SQS) ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	panic("ListQueues is not implemented")
}

// ListQueuesWithContext is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesWithContext(aws.Context, *sqs.ListQueuesInput, ...request.Option) (*sqs.ListQueuesOutput, error) {
	panic("ListQueuesWithContext is not implemented")
}

// ListQueuesRequest is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesRequest(*sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput) {
	panic("ListQueuesRequest is not implemented")
}
