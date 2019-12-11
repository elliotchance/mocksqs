package mocksqs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"time"
)

// SendMessage is partially supported. The following are not supported:
//
// - SendMessageInput.DelaySeconds
//
// - SendMessageInput.MessageAttributes
//
// - SendMessageInput.MessageDeduplicationId
//
// - SendMessageInput.MessageGroupId
//
// - SendMessageInput.MessageSystemAttributes
//
// - SendMessageOutput.MD5OfMessageAttributes
//
// - SendMessageOutput.MD5OfMessageBody
//
// - SendMessageOutput.MD5OfMessageSystemAttributes
//
// - SendMessageOutput.MessageId
//
// - SendMessageOutput.SequenceNumber
//
func (client *SQS) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	return client.sendMessage(input)
}

func (client *SQS) sendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	err := checkRequiredFields(map[string]interface{}{
		"SendMessageInput.QueueUrl":    input.QueueUrl,
		"SendMessageInput.MessageBody": input.MessageBody,
	})
	if err != nil {
		return &sqs.SendMessageOutput{}, err
	}

	if *input.MessageBody == "" {
		return &sqs.SendMessageOutput{}, errorMissingParameter("MessageBody")
	}

	if input.DelaySeconds != nil && (*input.DelaySeconds < 0 || *input.DelaySeconds > 900) {
		return &sqs.SendMessageOutput{}, errorInvalidParameterValue(
			fmt.Sprintf("Value %d for parameter DelaySeconds is invalid. Reason: DelaySeconds must be >= 0 and <= 900.",
				*input.DelaySeconds))
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		receiptHandle := uuid.New().String()
		queue.messages.Set(receiptHandle, &Message{
			Message: sqs.Message{
				Body:          input.MessageBody,
				ReceiptHandle: aws.String(receiptHandle),
			},
			VisibleAfter: time.Now(),
		})

		return &sqs.SendMessageOutput{}, nil
	}

	return &sqs.SendMessageOutput{}, errorNonExistentQueue()
}

// SendMessageWithContext is not implemented. It will panic in all cases.
func (client *SQS) SendMessageWithContext(aws.Context, *sqs.SendMessageInput, ...request.Option) (*sqs.SendMessageOutput, error) {
	panic("SendMessageWithContext is not implemented")
}

// SendMessageRequest is not implemented. It will panic in all cases.
func (client *SQS) SendMessageRequest(*sqs.SendMessageInput) (*request.Request, *sqs.SendMessageOutput) {
	panic("SendMessageRequest is not implemented")
}

// SendMessageBatch is not implemented. It will panic in all cases.
func (client *SQS) SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	panic("SendMessageBatch is not implemented")
}

// SendMessageBatchWithContext is not implemented. It will panic in all cases.
func (client *SQS) SendMessageBatchWithContext(aws.Context, *sqs.SendMessageBatchInput, ...request.Option) (*sqs.SendMessageBatchOutput, error) {
	panic("SendMessageBatchWithContext is not implemented")
}

// SendMessageBatchRequest is not implemented. It will panic in all cases.
func (client *SQS) SendMessageBatchRequest(*sqs.SendMessageBatchInput) (*request.Request, *sqs.SendMessageBatchOutput) {
	panic("SendMessageBatchRequest is not implemented")
}
