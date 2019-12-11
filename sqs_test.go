package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSQS(t *testing.T) {
	assert.Implements(t, (*sqsiface.SQSAPI)(nil), mocksqs.New())
}
