package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func assertSendMessageOutput(t *testing.T, expected, actual *sqs.SendMessageOutput) {
	assertAWSString(t, expected.MD5OfMessageAttributes, actual.MD5OfMessageAttributes)
	assertAWSString(t, expected.SequenceNumber, actual.SequenceNumber)
}

func TestSQS_SendMessage(t *testing.T) {
	client, url := getSQSClient()

	t.Run("MissingQueueURL", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			MessageBody: aws.String("a"),
		})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, SendMessageInput.QueueUrl.\n")
		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})

	t.Run("MissingMessageBody", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl: &url,
		})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, SendMessageInput.MessageBody.\n")
		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})

	t.Run("EmptyMessageBody", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    &url,
			MessageBody: aws.String(""),
		})
		assertRegexpError(t, err, `MissingParameter: The request must contain the parameter MessageBody.\s+status code: 400, request id: `+uuidRegexp+"$")
		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})

	t.Run("NegativeDelaySeconds", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:     &url,
			MessageBody:  aws.String("a"),
			DelaySeconds: aws.Int64(-1),
		})
		assertRegexpError(t, err, "InvalidParameterValue: Value -1 for parameter DelaySeconds is invalid. Reason: DelaySeconds must be >= 0 and <= 900.\n\tstatus code: 400, request id: "+uuidRegexp+"$")
		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})

	t.Run("DelaySecondsTooLarge", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:     &url,
			MessageBody:  aws.String("a"),
			DelaySeconds: aws.Int64(901),
		})
		assertRegexpError(t, err, "InvalidParameterValue: Value 901 for parameter DelaySeconds is invalid. Reason: DelaySeconds must be >= 0 and <= 900.\n\tstatus code: 400, request id: "+uuidRegexp+"$")
		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})

	t.Run("MinimalSuccess", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    &url,
			MessageBody: aws.String("a"),
		})
		require.NoError(t, err)

		assertSendMessageOutput(t, &sqs.SendMessageOutput{
			MD5OfMessageBody: aws.String("0cc175b9c0f1b6a831c399e269772661"),
		}, result)
	})

	t.Run("QueueDoesNotExist", func(t *testing.T) {
		result, err := client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    aws.String(url + "FOO"),
			MessageBody: aws.String("a"),
		})
		assertRegexpError(t, err, "AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.\n\tstatus code: 400, request id: "+uuidRegexp+"$")

		assertSendMessageOutput(t, &sqs.SendMessageOutput{}, result)
	})
}

func TestSQS_SendMessageWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "SendMessageWithContext is not implemented", func() {
		mocksqs.New().SendMessageWithContext(nil, nil)
	})
}

func TestSQS_SendMessageRequest(t *testing.T) {
	assert.PanicsWithValue(t, "SendMessageRequest is not implemented", func() {
		mocksqs.New().SendMessageRequest(nil)
	})
}

func TestSQS_SendMessageBatch(t *testing.T) {
	assert.PanicsWithValue(t, "SendMessageBatch is not implemented", func() {
		mocksqs.New().SendMessageBatch(nil)
	})
}

func TestSQS_SendMessageBatchWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "SendMessageBatchWithContext is not implemented", func() {
		mocksqs.New().SendMessageBatchWithContext(nil, nil)
	})
}

func TestSQS_SendMessageBatchRequest(t *testing.T) {
	assert.PanicsWithValue(t, "SendMessageBatchRequest is not implemented", func() {
		mocksqs.New().SendMessageBatchRequest(nil)
	})
}
