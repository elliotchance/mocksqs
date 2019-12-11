package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elliotchance/orderedmap"
	"sync"
	"time"
)

type Message struct {
	sqs.Message
	VisibleAfter time.Time
	ReceiveCount int64
}

type SQS struct {
	sync.RWMutex
	queues map[string]*orderedmap.OrderedMap
}

// New creates a new SQS service that contains no queues.
func New() *SQS {
	return &SQS{
		queues: make(map[string]*orderedmap.OrderedMap),
	}
}

// NewWithQueues create a new SQS service with prepopulated queues and messages.
//
// The map key is the queue URL. Each queue can contain zero or more messages.
// All of the messages on the queues will be immediately visible.
func NewWithQueues(queues map[string][]string) *SQS {
	client := New()

	for queueURL, messages := range queues {
		client.queues[queueURL] = orderedmap.NewOrderedMap()

		for _, body := range messages {
			_, _ = client.SendMessage(&sqs.SendMessageInput{
				QueueUrl:    aws.String(queueURL),
				MessageBody: aws.String(body),
			})
		}
	}

	return client
}
