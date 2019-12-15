package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQS_CreateQueue(t *testing.T) {
	t.Run("MissingQueueURL", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()
		t.Parallel()

		_, err := client.client.CreateQueue(&sqs.CreateQueueInput{})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, CreateQueueInput.QueueName.\n")
	})

	t.Run("QueueAlreadyExists", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()
		t.Parallel()

		result, err := client.client.CreateQueue(&sqs.CreateQueueInput{
			QueueName: aws.String(client.queueName),
		})
		require.NoError(t, err)
		assert.Equal(t, client.queueURL, *result.QueueUrl)
	})

	t.Run("Success", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()
		t.Parallel()

		queueName := uuid.New().String()
		result, err := client.client.CreateQueue(&sqs.CreateQueueInput{
			QueueName: &queueName,
		})
		require.NoError(t, err)

		defer func() {
			_, err := client.client.DeleteQueue(&sqs.DeleteQueueInput{
				QueueUrl: result.QueueUrl,
			})

			assert.NoError(t, err)
		}()

		assert.Contains(t, *result.QueueUrl, "/" + queueName)
	})
}

func TestSQS_CreateQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "CreateQueueWithContext is not implemented", func() {
		mocksqs.New().CreateQueueWithContext(nil, nil)
	})
}

func TestSQS_CreateQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "CreateQueueRequest is not implemented", func() {
		mocksqs.New().CreateQueueRequest(nil)
	})
}
