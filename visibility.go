package mocksqs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"time"
)

func (client *SQS) changeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		if message, ok := queue.messages.Get(*input.ReceiptHandle); ok {
			message.(*Message).VisibleAfter = time.Now().Add(
				time.Duration(*input.VisibilityTimeout) * time.Second)
		}

		return nil, nil
	}

	return nil, fmt.Errorf("no such queue: %s", *input.QueueUrl)
}

func (client *SQS) ChangeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	return client.changeMessageVisibility(input)
}

// ChangeMessageVisibilityWithContext is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityWithContext(aws.Context, *sqs.ChangeMessageVisibilityInput, ...request.Option) (*sqs.ChangeMessageVisibilityOutput, error) {
	panic("ChangeMessageVisibilityWithContext is not implemented")
}

// ChangeMessageVisibilityRequest is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) (*request.Request, *sqs.ChangeMessageVisibilityOutput) {
	panic("ChangeMessageVisibilityRequest is not implemented")
}

// ChangeMessageVisibilityBatch is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	panic("ChangeMessageVisibilityBatch is not implemented")
}

// ChangeMessageVisibilityBatchWithContext is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityBatchWithContext(aws.Context, *sqs.ChangeMessageVisibilityBatchInput, ...request.Option) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	panic("ChangeMessageVisibilityBatchWithContext is not implemented")
}

// ChangeMessageVisibilityBatchRequest is not implemented. It will panic in all cases.
func (client *SQS) ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) (*request.Request, *sqs.ChangeMessageVisibilityBatchOutput) {
	panic("ChangeMessageVisibilityBatchRequest is not implemented")
}
