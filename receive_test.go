package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSQS_ReceiveMessage(t *testing.T) {
	client, url := getSQSClient()

	t.Run("MissingQueueURL", func(t *testing.T) {
		_, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{})
		assert.EqualError(t, err, "InvalidParameter: 1 validation error(s) found.\n- missing required field, ReceiveMessageInput.QueueUrl.\n")
	})

	t.Run("QueueDoesNotExist", func(t *testing.T) {
		_, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(url + "FOO"),
		})
		assertRegexpError(t, err, "AWS.SimpleQueueService.NonExistentQueue: The specified queue does not exist for this wsdl version.\n\tstatus code: 400, request id: "+uuidRegexp+"$")
	})

	t.Run("ReceiveSingle", func(t *testing.T) {
		result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(url),
		})
		require.NoError(t, err)
		assert.Len(t, result.Messages, 1)
	})
}

func TestSQS_ReceiveMessageWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ReceiveMessageWithContext is not implemented", func() {
		mocksqs.New().ReceiveMessageWithContext(nil, nil)
	})
}

func TestSQS_ReceiveMessageRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ReceiveMessageRequest is not implemented", func() {
		mocksqs.New().ReceiveMessageRequest(nil)
	})
}
