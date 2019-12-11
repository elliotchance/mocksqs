package mocksqs

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"time"
)

// ReceiveMessage is partially supported. The following are not supported:
//
// - ReceiveMessageInput.AttributeNames
//
// - ReceiveMessageInput.MessageAttributeNames
//
// - ReceiveMessageInput.ReceiveRequestAttemptId
//
// - ReceiveMessageInput.VisibilityTimeout
//
// - ReceiveMessageInput.WaitTimeSeconds
//
func (client *SQS) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	client.Lock()
	defer client.Unlock()

	if input.QueueUrl == nil {
		return nil, errorMissingField("ReceiveMessageInput.QueueUrl")
	}

	if queue := client.GetQueue(*input.QueueUrl); queue != nil {
		output := &sqs.ReceiveMessageOutput{}

		for el := queue.messages.Front(); el != nil; el = el.Next() {
			message := el.Value.(*Message)
			if time.Now().After(message.VisibleAfter) {
				_, _ = client.changeMessageVisibility(&sqs.ChangeMessageVisibilityInput{
					QueueUrl:          input.QueueUrl,
					ReceiptHandle:     message.ReceiptHandle,
					VisibilityTimeout: aws.Int64(30),
				})

				message.ReceiveCount++

				output.Messages = append(output.Messages, &sqs.Message{
					Body:          message.Body,
					ReceiptHandle: message.ReceiptHandle,
					Attributes: map[string]*string{
						"ApproximateReceiveCount": aws.String(fmt.Sprintf("%d", message.ReceiveCount)),
					},
				})

				if input.MaxNumberOfMessages == nil {
					input.MaxNumberOfMessages = aws.Int64(1)
				}

				if len(output.Messages) == int(*input.MaxNumberOfMessages) {
					return output, nil
				}
			}
		}

		return output, nil
	}

	return nil, errorNonExistentQueue()
}

// ReceiveMessageWithContext is not implemented. It will panic in all cases.
func (client *SQS) ReceiveMessageWithContext(aws.Context, *sqs.ReceiveMessageInput, ...request.Option) (*sqs.ReceiveMessageOutput, error) {
	panic("ReceiveMessageWithContext is not implemented")
}

// ReceiveMessageRequest is not implemented. It will panic in all cases.
func (client *SQS) ReceiveMessageRequest(*sqs.ReceiveMessageInput) (*request.Request, *sqs.ReceiveMessageOutput) {
	panic("ReceiveMessageRequest is not implemented")
}
