package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/elliotchance/mocksqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const uuidRegexp = `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`

func assertRegexpError(t *testing.T, err error, regexp string) {
	require.Error(t, err)
	assert.Regexp(t, regexp, err.Error())
}

func assertAWSString(t *testing.T, expected, actual *string) {
	if expected == nil {
		assert.Nil(t, actual)
		return
	}

	require.NotNil(t, actual)
	assert.Equal(t, *expected, *actual)
}

func getRealSQSClient() (sqsiface.SQSAPI, string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sqs.New(sess), os.Getenv("QUEUE_URL")
}

func getMockSQSClient() (sqsiface.SQSAPI, string) {
	url := "https://sqs.us-east-1.amazonaws.com/281910179584/mocksqs"
	client := mocksqs.NewWithQueues(map[string][]string{
		url: {"foo"},
	})

	return client, url
}

func getSQSClient() (sqsiface.SQSAPI, string) {
	if os.Getenv("INTEGRATION") != "" {
		return getRealSQSClient()
	}

	return getMockSQSClient()
}
