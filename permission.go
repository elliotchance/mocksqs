package mocksqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// AddPermission is not implemented. It will panic in all cases.
func (client *SQS) AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	panic("AddPermission is not implemented")
}

// AddPermissionWithContext is not implemented. It will panic in all cases.
func (client *SQS) AddPermissionWithContext(aws.Context, *sqs.AddPermissionInput, ...request.Option) (*sqs.AddPermissionOutput, error) {
	panic("AddPermissionWithContext is not implemented")
}

// AddPermissionRequest is not implemented. It will panic in all cases.
func (client *SQS) AddPermissionRequest(*sqs.AddPermissionInput) (*request.Request, *sqs.AddPermissionOutput) {
	panic("AddPermissionRequest is not implemented")
}

// RemovePermission is not implemented. It will panic in all cases.
func (client *SQS) RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	panic("RemovePermission is not implemented")
}

// RemovePermissionWithContext is not implemented. It will panic in all cases.
func (client *SQS) RemovePermissionWithContext(aws.Context, *sqs.RemovePermissionInput, ...request.Option) (*sqs.RemovePermissionOutput, error) {
	panic("RemovePermissionWithContext is not implemented")
}

// RemovePermissionRequest is not implemented. It will panic in all cases.
func (client *SQS) RemovePermissionRequest(*sqs.RemovePermissionInput) (*request.Request, *sqs.RemovePermissionOutput) {
	panic("RemovePermissionRequest is not implemented")
}
