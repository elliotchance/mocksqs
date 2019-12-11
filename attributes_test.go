package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_GetQueueAttributes(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueAttributes is not implemented", func() {
		mocksqs.New().GetQueueAttributes(nil)
	})
}

func TestSQS_GetQueueAttributesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueAttributesWithContext is not implemented", func() {
		mocksqs.New().GetQueueAttributesWithContext(nil, nil)
	})
}

func TestSQS_GetQueueAttributesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueAttributesRequest is not implemented", func() {
		mocksqs.New().GetQueueAttributesRequest(nil)
	})
}

func TestSQS_SetQueueAttributes(t *testing.T) {
	assert.PanicsWithValue(t, "SetQueueAttributes is not implemented", func() {
		mocksqs.New().SetQueueAttributes(nil)
	})
}

func TestSQS_SetQueueAttributesWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "SetQueueAttributesWithContext is not implemented", func() {
		mocksqs.New().SetQueueAttributesWithContext(nil, nil)
	})
}

func TestSQS_SetQueueAttributesRequest(t *testing.T) {
	assert.PanicsWithValue(t, "SetQueueAttributesRequest is not implemented", func() {
		mocksqs.New().SetQueueAttributesRequest(nil)
	})
}
