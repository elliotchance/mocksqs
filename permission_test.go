package mocksqs_test

import (
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQS_AddPermission(t *testing.T) {
	assert.PanicsWithValue(t, "AddPermission is not implemented", func() {
		mocksqs.New().AddPermission(nil)
	})
}

func TestSQS_AddPermissionWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "AddPermissionWithContext is not implemented", func() {
		mocksqs.New().AddPermissionWithContext(nil, nil)
	})
}

func TestSQS_AddPermissionRequest(t *testing.T) {
	assert.PanicsWithValue(t, "AddPermissionRequest is not implemented", func() {
		mocksqs.New().AddPermissionRequest(nil)
	})
}

func TestSQS_RemovePermission(t *testing.T) {
	assert.PanicsWithValue(t, "RemovePermission is not implemented", func() {
		mocksqs.New().RemovePermission(nil)
	})
}

func TestSQS_RemovePermissionWithContext(t *testing.T) {
	assert.PanicsWithValue(t, "RemovePermissionWithContext is not implemented", func() {
		mocksqs.New().RemovePermissionWithContext(nil, nil)
	})
}

func TestSQS_RemovePermissionRequest(t *testing.T) {
	assert.PanicsWithValue(t, "RemovePermissionRequest is not implemented", func() {
		mocksqs.New().RemovePermissionRequest(nil)
	})
}
