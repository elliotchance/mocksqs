package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_ListDeadLetterSourceQueues(t *testing.T) {
	assert.PanicsWithValue(t, "ListDeadLetterSourceQueues is not implemented", func() {
		mocksqs.New().ListDeadLetterSourceQueues(nil)
	})
}

func TestSQS_ListDeadLetterSourceQueuesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListDeadLetterSourceQueuesWithContext is not implemented", func() {
		mocksqs.New().ListDeadLetterSourceQueuesWithContext(nil, nil)
	})
}

func TestSQS_ListDeadLetterSourceQueuesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListDeadLetterSourceQueuesRequest is not implemented", func() {
		mocksqs.New().ListDeadLetterSourceQueuesRequest(nil)
	})
}
