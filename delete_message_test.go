package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQS_DeleteMessage(t *testing.T) {
	t.Run("MissingQueueURL", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()

		_, err := client.client.DeleteMessage(&sqs.DeleteMessageInput{
			ReceiptHandle: aws.String("a"),
		})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, DeleteMessageInput.QueueUrl.\n")
	})

	t.Run("MissingReceiptHandle", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()

		_, err := client.client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl: &client.queueURL,
		})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, DeleteMessageInput.ReceiptHandle.\n")
	})

	t.Run("ReceiptHandleDoesNotExist", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()

		_, err := client.client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &client.queueURL,
			ReceiptHandle: aws.String("MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+CwLj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QEauMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0="),
		})
		assertRegexpError(t, err, `InternalError: We encountered an internal error. Please try again.\s+status code: [45]00, request id: `+uuidRegexp+"$")
	})

	t.Run("ReceiptHandleDoesExist", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()

		_, err := client.client.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    &client.queueURL,
			MessageBody: aws.String("a"),
		})
		require.NoError(t, err)

		result1, err := client.client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: &client.queueURL,
		})
		require.NoError(t, err)
		require.Len(t, result1.Messages, 1)

		_, err = client.client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &client.queueURL,
			ReceiptHandle: result1.Messages[0].ReceiptHandle,
		})
		require.NoError(t, err)
	})

	t.Run("QueueDoesNotExist", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()

		_, err := client.client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      aws.String(mocksqs.CreateQueueURL("foo")),
			ReceiptHandle: aws.String("MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+CwLj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QEauMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0="),
		})
		assertRegexpError(t, err, "AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.\n\tstatus code: 400, request id: "+uuidRegexp+"$")
	})
}

func TestSQS_DeleteMessageWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteMessageWithContext is not implemented", func() {
		mocksqs.New().DeleteMessageWithContext(nil, nil)
	})
}

func TestSQS_DeleteMessageRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteMessageRequest is not implemented", func() {
		mocksqs.New().DeleteMessageRequest(nil)
	})
}

func TestSQS_DeleteMessageBatchWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteMessageBatchWithContext is not implemented", func() {
		mocksqs.New().DeleteMessageBatchWithContext(nil, nil)
	})
}

func TestSQS_DeleteMessageBatchRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteMessageBatchRequest is not implemented", func() {
		mocksqs.New().DeleteMessageBatchRequest(nil)
	})
}
