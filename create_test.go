package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_CreateQueue(t *testing.T) {
	assert.PanicsWithValue(t, "CreateQueue is not implemented", func() {
		mocksqs.New().CreateQueue(nil)
	})
}

func TestSQS_CreateQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "CreateQueueWithContext is not implemented", func() {
		mocksqs.New().CreateQueueWithContext(nil, nil)
	})
}

func TestSQS_CreateQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "CreateQueueRequest is not implemented", func() {
		mocksqs.New().CreateQueueRequest(nil)
	})
}
