package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"math/rand"
	"sync"
	"time"
)

type SQS struct {
	sync.RWMutex
	queues sync.Map // map[string]*Queue

	// SimulateHTTPLatency when enabled will add a sleep between 20 and 100
	// milliseconds to each call that would otherwise need to make a HTTP
	// request with a real SQS client.
	SimulateHTTPLatency bool
}

// New creates a new SQS service that contains no queues.
func New() *SQS {
	return &SQS{
		queues: sync.Map{},
	}
}

// NewWithQueues create a new SQS service with prepopulated queues and messages.
//
// The map key is the queue URL. Each queue can contain zero or more messages.
// All of the messages on the queues will be immediately visible.
func NewWithQueues(queues map[string][]string) *SQS {
	client := New()

	for queueURL, messages := range queues {
		client.queues.Store(queueURL, newQueue(queueURL))

		for _, body := range messages {
			_, _ = client.SendMessage(&sqs.SendMessageInput{
				QueueUrl:    aws.String(queueURL),
				MessageBody: aws.String(body),
			})
		}
	}

	return client
}

func (client *SQS) GetQueue(queueURL string) *Queue {
	queue, ok := client.queues.Load(queueURL)
	if !ok {
		return nil
	}

	return queue.(*Queue)
}

func (client *SQS) httpRequest() {
	if client.SimulateHTTPLatency {
		ms := 20 + rand.Int()%80
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}
