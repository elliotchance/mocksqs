package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_PurgeQueue(t *testing.T) {
	assert.PanicsWithValue(t, "PurgeQueue is not implemented", func() {
		mocksqs.New().PurgeQueue(nil)
	})
}

func TestSQS_PurgeQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "PurgeQueueWithContext is not implemented", func() {
		mocksqs.New().PurgeQueueWithContext(nil, nil)
	})
}

func TestSQS_PurgeQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "PurgeQueueRequest is not implemented", func() {
		mocksqs.New().PurgeQueueRequest(nil)
	})
}
