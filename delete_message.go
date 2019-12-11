package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// DeleteMessage is fully supported.
func (client *SQS) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	client.Lock()
	defer client.Unlock()

	return client.deleteMessage(input)
}

func (client *SQS) deleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"DeleteMessageInput.QueueUrl":      input.QueueUrl,
		"DeleteMessageInput.ReceiptHandle": input.ReceiptHandle,
	})
	if err != nil {
		return nil, err
	}

	if queue, ok := client.queues[*input.QueueUrl]; ok {
		didDelete := queue.Delete(*input.ReceiptHandle)
		if !didDelete {
			return nil, errorInternal()
		}

		return nil, nil
	}

	return nil, errorNonExistentQueue()
}

// DeleteMessageBatch is fully supported.
func (client *SQS) DeleteMessageBatch(input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	client.Lock()
	defer client.Unlock()

	for _, message := range input.Entries {
		_, err := client.deleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      input.QueueUrl,
			ReceiptHandle: message.ReceiptHandle,
		})

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// DeleteMessageWithContext is not implemented. It will panic in all cases.
func (client *SQS) DeleteMessageWithContext(aws.Context, *sqs.DeleteMessageInput, ...request.Option) (*sqs.DeleteMessageOutput, error) {
	panic("DeleteMessageWithContext is not implemented")
}

// DeleteMessageRequest is not implemented. It will panic in all cases.
func (client *SQS) DeleteMessageRequest(*sqs.DeleteMessageInput) (*request.Request, *sqs.DeleteMessageOutput) {
	panic("DeleteMessageRequest is not implemented")
}

// DeleteMessageBatchWithContext is not implemented. It will panic in all cases.
func (client *SQS) DeleteMessageBatchWithContext(aws.Context, *sqs.DeleteMessageBatchInput, ...request.Option) (*sqs.DeleteMessageBatchOutput, error) {
	panic("DeleteMessageBatchWithContext is not implemented")
}

// DeleteMessageBatchRequest is not implemented. It will panic in all cases.
func (client *SQS) DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) (*request.Request, *sqs.DeleteMessageBatchOutput) {
	panic("DeleteMessageBatchRequest is not implemented")
}
