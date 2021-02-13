package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"strings"
)

// ListQueues is fully supported.
func (client *SQS) ListQueues(input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	client.httpRequest()

	prefix := ""
	if input.QueueNamePrefix != nil {
		prefix = *input.QueueNamePrefix
	}

	output := &sqs.ListQueuesOutput{}

	client.queues.Range(func(key, value interface{}) bool {
		if strings.HasPrefix(key.(string), prefix) {
			output.QueueUrls = append(output.QueueUrls, &value.(*Queue).URL)
		}

		return true
	})

	return output, nil
}

// ListQueuesWithContext is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesWithContext(aws.Context, *sqs.ListQueuesInput, ...request.Option) (*sqs.ListQueuesOutput, error) {
	panic("ListQueuesWithContext is not implemented")
}

// ListQueuesRequest is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesRequest(*sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput) {
	panic("ListQueuesRequest is not implemented")
}

// ListQueuesPages is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesPages(*sqs.ListQueuesInput, func(*sqs.ListQueuesOutput, bool) bool) error {
	panic("ListQueuesPages is not implemented")
}

// ListQueuesPagesWithContext is not implemented. It will panic in all cases.
func (client *SQS) ListQueuesPagesWithContext(aws.Context, *sqs.ListQueuesInput, func(*sqs.ListQueuesOutput, bool) bool, ...request.Option) error {
	panic("ListQueuesPagesWithContext is not implemented")
}
