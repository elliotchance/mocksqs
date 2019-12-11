package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueAttributes is not implemented. It will panic in all cases.
func (client *SQS) GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	panic("GetQueueAttributes is not implemented")
}

// GetQueueAttributesWithContext is not implemented. It will panic in all cases.
func (client *SQS) GetQueueAttributesWithContext(aws.Context, *sqs.GetQueueAttributesInput, ...request.Option) (*sqs.GetQueueAttributesOutput, error) {
	panic("GetQueueAttributesWithContext is not implemented")
}

// GetQueueAttributesRequest is not implemented. It will panic in all cases.
func (client *SQS) GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) (*request.Request, *sqs.GetQueueAttributesOutput) {
	panic("GetQueueAttributesRequest is not implemented")
}

// SetQueueAttributes is not implemented. It will panic in all cases.
func (client *SQS) SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	panic("SetQueueAttributes is not implemented")
}

// SetQueueAttributesWithContext is not implemented. It will panic in all cases.
func (client *SQS) SetQueueAttributesWithContext(aws.Context, *sqs.SetQueueAttributesInput, ...request.Option) (*sqs.SetQueueAttributesOutput, error) {
	panic("SetQueueAttributesWithContext is not implemented")
}

// SetQueueAttributesRequest is not implemented. It will panic in all cases.
func (client *SQS) SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) (*request.Request, *sqs.SetQueueAttributesOutput) {
	panic("SetQueueAttributesRequest is not implemented")
}
