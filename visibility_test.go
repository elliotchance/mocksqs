package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_ChangeMessageVisibilityWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ChangeMessageVisibilityWithContext is not implemented", func() {
		mocksqs.New().ChangeMessageVisibilityWithContext(nil, nil)
	})
}

func TestSQS_ChangeMessageVisibilityRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ChangeMessageVisibilityRequest is not implemented", func() {
		mocksqs.New().ChangeMessageVisibilityRequest(nil)
	})
}

func TestSQS_ChangeMessageVisibilityBatch(t *testing.T) {
	assert.PanicsWithValue(t, "ChangeMessageVisibilityBatch is not implemented", func() {
		mocksqs.New().ChangeMessageVisibilityBatch(nil)
	})
}

func TestSQS_ChangeMessageVisibilityBatchWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ChangeMessageVisibilityBatchWithContext is not implemented", func() {
		mocksqs.New().ChangeMessageVisibilityBatchWithContext(nil, nil)
	})
}

func TestSQS_ChangeMessageVisibilityBatchRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ChangeMessageVisibilityBatchRequest is not implemented", func() {
		mocksqs.New().ChangeMessageVisibilityBatchRequest(nil)
	})
}
