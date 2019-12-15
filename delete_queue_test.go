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

func TestSQS_DeleteQueue(t *testing.T) {
	t.Run("MissingQueueURL", func(t *testing.T) {
		client := getSQSClient()
		defer client.cleanup()
		t.Parallel()

		_, err := client.client.DeleteQueue(&sqs.DeleteQueueInput{})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, DeleteQueueInput.QueueUrl.\n")
	})

	t.Run("QueueDoesNotExist", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()
		t.Parallel()

		_, err := client.client.DeleteQueue(&sqs.DeleteQueueInput{
			QueueUrl: aws.String(client.queueURL + "foo"),
		})
		assertRegexpError(t, err, `AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.\s+status code: [45]00, request id: `+uuidRegexp+"$")
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

		_, err = client.client.DeleteQueue(&sqs.DeleteQueueInput{
			QueueUrl: result.QueueUrl,
		})
		assert.NoError(t, err)
	})
}

func TestSQS_DeleteQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteQueueWithContext is not implemented", func() {
		mocksqs.New().DeleteQueueWithContext(nil, nil)
	})
}

func TestSQS_DeleteQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteQueueRequest is not implemented", func() {
		mocksqs.New().DeleteQueueRequest(nil)
	})
}
