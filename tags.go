package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// ListQueueTags is not implemented. It will panic in all cases.
func (client *SQS) ListQueueTags(*sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	panic("ListQueueTags is not implemented")
}

// ListQueueTagsWithContext is not implemented. It will panic in all cases.
func (client *SQS) ListQueueTagsWithContext(aws.Context, *sqs.ListQueueTagsInput, ...request.Option) (*sqs.ListQueueTagsOutput, error) {
	panic("ListQueueTagsWithContext is not implemented")
}

// ListQueueTagsRequest is not implemented. It will panic in all cases.
func (client *SQS) ListQueueTagsRequest(*sqs.ListQueueTagsInput) (*request.Request, *sqs.ListQueueTagsOutput) {
	panic("ListQueueTagsRequest is not implemented")
}

// TagQueue is not implemented. It will panic in all cases.
func (client *SQS) TagQueue(*sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	panic("TagQueue is not implemented")
}

// TagQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) TagQueueWithContext(aws.Context, *sqs.TagQueueInput, ...request.Option) (*sqs.TagQueueOutput, error) {
	panic("TagQueueWithContext is not implemented")
}

// TagQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) TagQueueRequest(*sqs.TagQueueInput) (*request.Request, *sqs.TagQueueOutput) {
	panic("TagQueueRequest is not implemented")
}

// UntagQueue is not implemented. It will panic in all cases.
func (client *SQS) UntagQueue(*sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	panic("UntagQueue is not implemented")
}

// UntagQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) UntagQueueWithContext(aws.Context, *sqs.UntagQueueInput, ...request.Option) (*sqs.UntagQueueOutput, error) {
	panic("UntagQueueWithContext is not implemented")
}

// UntagQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) UntagQueueRequest(*sqs.UntagQueueInput) (*request.Request, *sqs.UntagQueueOutput) {
	panic("UntagQueueRequest is not implemented")
}
