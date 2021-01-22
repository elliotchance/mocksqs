package mocksqs

import (
	"github.com/elliotchance/orderedmap"
)

type Queue struct {
	messages *orderedmap.OrderedMap

	URL string

	// OnEmptyQueue is triggered after the last message is deleted from the
	// queue. If a queue becomes empty multiple times then this is triggered
	// each time.
	OnEmptyQueue func()
}

func newQueue(url string) *Queue {
	return &Queue{
		messages: orderedmap.NewOrderedMap(),
		URL:      url,
	}
}

func (queue *Queue) delete(receiptHandle string) (didDelete bool) {
	didDelete = queue.messages.Delete(receiptHandle)

	if queue.OnEmptyQueue != nil && queue.messages.Len() == 0 {
		queue.OnEmptyQueue()
	}

	return
}

func (queue *Queue) Messages() (*orderedmap.OrderedMap, bool) {
	m := queue.messages

	rv := orderedmap.NewOrderedMap()
	for _, key := range m.Keys() {
		v, ok := m.Get(key)
		if !ok {
			return nil, false
		}
		rv.Set(key, v)
	}
	return rv, true
}
