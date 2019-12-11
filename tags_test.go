package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_ListQueueTags(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueueTags is not implemented", func() {
		mocksqs.New().ListQueueTags(nil)
	})
}

func TestSQS_ListQueueTagsWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueueTagsWithContext is not implemented", func() {
		mocksqs.New().ListQueueTagsWithContext(nil, nil)
	})
}

func TestSQS_ListQueueTagsRequest(t *testing.T) {
	assert.PanicsWithValue(t, "ListQueueTagsRequest is not implemented", func() {
		mocksqs.New().ListQueueTagsRequest(nil)
	})
}

func TestSQS_TagQueue(t *testing.T) {
	assert.PanicsWithValue(t, "TagQueue is not implemented", func() {
		mocksqs.New().TagQueue(nil)
	})
}

func TestSQS_TagQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "TagQueueWithContext is not implemented", func() {
		mocksqs.New().TagQueueWithContext(nil, nil)
	})
}

func TestSQS_TagQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "TagQueueRequest is not implemented", func() {
		mocksqs.New().TagQueueRequest(nil)
	})
}

func TestSQS_UntagQueue(t *testing.T) {
	assert.PanicsWithValue(t, "UntagQueue is not implemented", func() {
		mocksqs.New().UntagQueue(nil)
	})
}

func TestSQS_UntagQueueWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "UntagQueueWithContext is not implemented", func() {
		mocksqs.New().UntagQueueWithContext(nil, nil)
	})
}

func TestSQS_UntagQueueRequest(t *testing.T) {
	assert.PanicsWithValue(t, "UntagQueueRequest is not implemented", func() {
		mocksqs.New().UntagQueueRequest(nil)
	})
}
