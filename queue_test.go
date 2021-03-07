package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueue_Delete(t *testing.T) {
	queueURL := "queueURL"

	t.Run("QueueDidBecomeEmpty", func(t *testing.T) {
		client := mocksqs.NewWithQueues(map[string][]string{
			queueURL: {"foo"},
		})

		receive, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: &queueURL,
		})
		require.NoError(t, err)

		_, err = client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &queueURL,
			ReceiptHandle: receive.Messages[0].ReceiptHandle,
		})
		require.NoError(t, err)
	})

	t.Run("QueueDidBecomeEmptyWithFunction", func(t *testing.T) {
		didCall := false
		client := mocksqs.NewWithQueues(map[string][]string{
			queueURL: {"foo"},
		})
		client.GetQueue(queueURL).OnEmptyQueue = func() {
			didCall = true
		}

		receive, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: &queueURL,
		})
		require.NoError(t, err)

		_, err = client.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      &queueURL,
			ReceiptHandle: receive.Messages[0].ReceiptHandle,
		})
		require.NoError(t, err)

		assert.True(t, didCall)
	})
}

func TestQueue_Messages(t *testing.T) {
	queueURL := "queueURL"

	t.Run("QueueDumpsMessages", func(t *testing.T) {
		client := mocksqs.NewWithQueues(map[string][]string{
			queueURL: {},
		})

		mb := "messageBody"

		_, err := client.SendMessage(&sqs.SendMessageInput{
			MessageBody:             &mb,
			QueueUrl:                &queueURL,
		})

		m := client.GetQueue(queueURL).Messages()

		require.NoError(t, err)
		assert.Equal(t, 1, m.Len())
	})

	t.Run("SendingMessageAfterCopyDoesNotChangeCopy", func(t *testing.T) {
		client := mocksqs.NewWithQueues(map[string][]string{
			queueURL: {},
		})

		mb := "messageBody"

		_, err1 := client.SendMessage(&sqs.SendMessageInput{
			MessageBody:             &mb,
			QueueUrl:                &queueURL,
		})
		require.NoError(t, err1)

		m := client.GetQueue(queueURL).Messages()

		_, err2 := client.SendMessage(&sqs.SendMessageInput{
			MessageBody:             &mb,
			QueueUrl:                &queueURL,
		})

		require.NoError(t, err2)
		assert.Equal(t, 1, m.Len())
	})
}
