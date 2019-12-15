package mocksqs_test

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/elliotchance/mocksqs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const uuidRegexp = `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`

type clientDetails struct {
	client    sqsiface.SQSAPI
	queueURL  string
	queueName string
	cleanup   func()
}

func assertRegexpError(t *testing.T, err error, regexp string) {
	require.Error(t, err)
	assert.Regexp(t, regexp, err.Error())
}

func dereferenceAWSStrings(ss1 []*string) (ss2 []string) {
	for _, ss := range ss1 {
		ss2 = append(ss2, *ss)
	}

	return
}

func assertContainsAWSString(t *testing.T, expected string, actual []*string) {
	assert.Contains(t, dereferenceAWSStrings(actual), expected)
}

func assertAWSString(t *testing.T, expected, actual *string) {
	if expected == nil {
		assert.Nil(t, actual)
		return
	}

	require.NotNil(t, actual)
	assert.Equal(t, *expected, *actual)
}

func getRealSQSClient() *clientDetails {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &clientDetails{
		client:  sqs.New(sess),
		cleanup: func() {},
	}
}

func getMockSQSClient() *clientDetails {
	client := mocksqs.New()

	return &clientDetails{
		client:  client,
		cleanup: func() {},
	}
}

func getSQSClient() *clientDetails {
	if os.Getenv("INTEGRATION") != "" {
		return getRealSQSClient()
	}

	return getMockSQSClient()
}

func getSQSClientWithQueue() *clientDetails {
	client := getSQSClient()
	client.queueName = uuid.New().String()

	result, err := client.client.CreateQueue(&sqs.CreateQueueInput{
		QueueName: &client.queueName,
	})
	if err != nil {
		panic(err)
	}
	client.queueURL = *result.QueueUrl

	client.cleanup = func() {
		_, err := client.client.DeleteQueue(&sqs.DeleteQueueInput{
			QueueUrl: &client.queueURL,
		})
		if err != nil {
			panic(err)
		}
	}

	return client
}
