package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_ListQueues(t *testing.T) {
	t.Run("WithoutQueueNamePrefix", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()

		result, err := client.client.ListQueues(&sqs.ListQueuesInput{})
		assert.NoError(t, err)
		assertContainsAWSString(t, client.queueURL, result.QueueUrls)
	})

	t.Run("WithQueueNamePrefix", func(t *testing.T) {
		client := getSQSClientWithQueue()
		defer client.cleanup()

		result, err := client.client.ListQueues(&sqs.ListQueuesInput{
			QueueNamePrefix: aws.String("wont-match-anything"),
		})
		assert.NoError(t, err)
		assert.Len(t, result.QueueUrls, 0)
	})
}

func TestSQS_ListQueuesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesWithContext is not implemented", func() {
		mocksqs.New().ListQueuesWithContext(nil, nil)
	})
}

func TestSQS_ListQueuesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesRequest is not implemented", func() {
		mocksqs.New().ListQueuesRequest(nil)
	})
}

func TestSQS_ListQueuesPages(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesPages is not implemented", func() {
		mocksqs.New().ListQueuesPages(nil, nil)
	})
}

func TestSQS_ListQueuesPagesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesPagesWithContext is not implemented", func() {
		mocksqs.New().ListQueuesPagesWithContext(nil, nil, nil)
	})
}
