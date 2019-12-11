package mocksqs

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"time"
)

type Message struct {
	sqs.Message
	VisibleAfter time.Time
	ReceiveCount int64
}
