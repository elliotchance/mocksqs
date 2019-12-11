package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_ListQueues(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueues is not implemented", func() {
		mocksqs.New().ListQueues(nil)
	})
}

func TestSQS_ListQueuesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesWithContext is not implemented", func() {
		mocksqs.New().ListQueuesWithContext(nil, nil)
	})
}

func TestSQS_ListQueuesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueuesRequest is not implemented", func() {
		mocksqs.New().ListQueuesRequest(nil)
	})
}
