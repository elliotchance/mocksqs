package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_DeleteQueue(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteQueue is not implemented", func() {
		mocksqs.New().DeleteQueue(nil)
	})
}

func TestSQS_DeleteQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteQueueWithContext is not implemented", func() {
		mocksqs.New().DeleteQueueWithContext(nil, nil)
	})
}

func TestSQS_DeleteQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "DeleteQueueRequest is not implemented", func() {
		mocksqs.New().DeleteQueueRequest(nil)
	})
}
