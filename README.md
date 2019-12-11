# 📤 github.com/elliotchance/mocksqs

[![Build Status](https://travis-ci.org/elliotchance/mocksqs.svg?branch=master)](https://travis-ci.org/elliotchance/mocksqs)
[![GoDoc](https://godoc.org/github.com/elliotchance/mocksqs?status.svg)](https://godoc.org/github.com/elliotchance/mocksqs)

# Creating the Service

The simplest way to create a new SQS service is with `mocksqs.New()`. However,
if you need queues prepopulated you can use `mocksqs.NewWithQueues()`:

```go
url := "https://sqs.us-east-1.amazonaws.com/281910179584/mocksqs"
client := mocksqs.NewWithQueues(map[string][]string{
	url: {"foo", "bar"},
})

result, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
	QueueUrl: aws.String(url),
})
```

# Supported Functionality

Only some of the common SQS methods are implemented. Methods not implemented
will panic.

You can view the specific implementation details in the
[godoc documentation](https://godoc.org/github.com/elliotchance/mocksqs).
