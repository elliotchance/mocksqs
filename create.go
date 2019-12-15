package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// CreateQueue is partially implemented. The following fields are not
// implemented:
//
// - Attributes: DelaySeconds, MaximumMessageSize, MessageRetentionPeriod,
// Policy, ReceiveMessageWaitTimeSeconds, RedrivePolicy, VisibilityTimeout,
// KmsMasterKeyId, KmsDataKeyReusePeriodSeconds, FifoQueue,
// ContentBasedDeduplication
//
// - Tags
//
func (client *SQS) CreateQueue(input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	client.httpRequest()

	client.Lock()
	defer client.Unlock()

	if input.QueueName == nil {
		return &sqs.CreateQueueOutput{},
			errorMissingField("CreateQueueInput.QueueName")
	}

	queueURL := CreateQueueURL(*input.QueueName)

	if client.GetQueue(queueURL) == nil {
		client.queues.Store(queueURL, newQueue())
	}

	return &sqs.CreateQueueOutput{
		QueueUrl: &queueURL,
	}, nil
}

// CreateQueueWithContext is not implemented. It will panic in all cases.
func (client *SQS) CreateQueueWithContext(aws.Context, *sqs.CreateQueueInput, ...request.Option) (*sqs.CreateQueueOutput, error) {
	panic("CreateQueueWithContext is not implemented")
}

// CreateQueueRequest is not implemented. It will panic in all cases.
func (client *SQS) CreateQueueRequest(*sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput) {
	panic("CreateQueueRequest is not implemented")
}
