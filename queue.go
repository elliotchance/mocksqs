package mocksqs

import (
	"github.com/elliotchance/orderedmap"
)

type Queue struct {
	messages *orderedmap.OrderedMap

	// OnEmptyQueue is triggered after the last message is deleted from the
	// queue. If a queue becomes empty multiple times then this is triggered
	// each time.
	OnEmptyQueue func()
}

func newQueue() *Queue {
	return &Queue{
		messages: orderedmap.NewOrderedMap(),
	}
}

func (queue *Queue) delete(receiptHandle string) (didDelete bool) {
	didDelete = queue.messages.Delete(receiptHandle)

	if queue.OnEmptyQueue != nil && queue.messages.Len() == 0 {
		queue.OnEmptyQueue()
	}

	return
}
