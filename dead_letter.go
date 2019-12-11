package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ListDeadLetterSourceQueues is not implemented. It will panic in all cases.
func (client *SQS) ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	panic("ListDeadLetterSourceQueues is not implemented")
}

// ListDeadLetterSourceQueuesWithContext is not implemented. It will panic in all cases.
func (client *SQS) ListDeadLetterSourceQueuesWithContext(aws.Context, *sqs.ListDeadLetterSourceQueuesInput, ...request.Option) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	panic("ListDeadLetterSourceQueuesWithContext is not implemented")
}

// ListDeadLetterSourceQueuesRequest is not implemented. It will panic in all cases.
func (client *SQS) ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) (*request.Request, *sqs.ListDeadLetterSourceQueuesOutput) {
	panic("ListDeadLetterSourceQueuesRequest is not implemented")
}

