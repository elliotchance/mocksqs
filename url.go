package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueUrl is not implemented. It will panic in all cases.
func (client *SQS) GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	panic("GetQueueUrl is not implemented")
}

// GetQueueUrlWithContext is not implemented. It will panic in all cases.
func (client *SQS) GetQueueUrlWithContext(aws.Context, *sqs.GetQueueUrlInput, ...request.Option) (*sqs.GetQueueUrlOutput, error) {
	panic("GetQueueUrlWithContext is not implemented")
}

// GetQueueUrlRequest is not implemented. It will panic in all cases.
func (client *SQS) GetQueueUrlRequest(*sqs.GetQueueUrlInput) (*request.Request, *sqs.GetQueueUrlOutput) {
	panic("GetQueueUrlRequest is not implemented")
}
