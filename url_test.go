package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_GetQueueUrl(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueUrl is not implemented", func() {
		mocksqs.New().GetQueueUrl(nil)
	})
}

func TestSQS_GetQueueUrlWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueUrlWithContext is not implemented", func() {
		mocksqs.New().GetQueueUrlWithContext(nil, nil)
	})
}

func TestSQS_GetQueueUrlRequest(t *testing.T) {
	assert.PanicsWithValue(t, "GetQueueUrlRequest is not implemented", func() {
		mocksqs.New().GetQueueUrlRequest(nil)
	})
}
