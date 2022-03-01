// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAttachmentIdNotFound for service response error code
	// "AttachmentIdNotFound".
	//
	// An attachment with the specified ID could not be found.
	ErrCodeAttachmentIdNotFound = "AttachmentIdNotFound"

	// ErrCodeAttachmentLimitExceeded for service response error code
	// "AttachmentLimitExceeded".
	//
	// The limit for the number of attachment sets created in a short period of
	// time has been exceeded.
	ErrCodeAttachmentLimitExceeded = "AttachmentLimitExceeded"

	// ErrCodeAttachmentSetExpired for service response error code
	// "AttachmentSetExpired".
	//
	// The expiration time of the attachment set has passed. The set expires 1 hour
	// after it is created.
	ErrCodeAttachmentSetExpired = "AttachmentSetExpired"

	// ErrCodeAttachmentSetIdNotFound for service response error code
	// "AttachmentSetIdNotFound".
	//
	// An attachment set with the specified ID could not be found.
	ErrCodeAttachmentSetIdNotFound = "AttachmentSetIdNotFound"

	// ErrCodeAttachmentSetSizeLimitExceeded for service response error code
	// "AttachmentSetSizeLimitExceeded".
	//
	// A limit for the size of an attachment set has been exceeded. The limits are
	// three attachments and 5 MB per attachment.
	ErrCodeAttachmentSetSizeLimitExceeded = "AttachmentSetSizeLimitExceeded"

	// ErrCodeCaseCreationLimitExceeded for service response error code
	// "CaseCreationLimitExceeded".
	//
	// The case creation limit for the account has been exceeded.
	ErrCodeCaseCreationLimitExceeded = "CaseCreationLimitExceeded"

	// ErrCodeCaseIdNotFound for service response error code
	// "CaseIdNotFound".
	//
	// The requested caseId couldn't be located.
	ErrCodeCaseIdNotFound = "CaseIdNotFound"

	// ErrCodeDescribeAttachmentLimitExceeded for service response error code
	// "DescribeAttachmentLimitExceeded".
	//
	// The limit for the number of DescribeAttachment requests in a short period
	// of time has been exceeded.
	ErrCodeDescribeAttachmentLimitExceeded = "DescribeAttachmentLimitExceeded"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// An internal server error occurred.
	ErrCodeInternalServerError = "InternalServerError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AttachmentIdNotFound":            newErrorAttachmentIdNotFound,
	"AttachmentLimitExceeded":         newErrorAttachmentLimitExceeded,
	"AttachmentSetExpired":            newErrorAttachmentSetExpired,
	"AttachmentSetIdNotFound":         newErrorAttachmentSetIdNotFound,
	"AttachmentSetSizeLimitExceeded":  newErrorAttachmentSetSizeLimitExceeded,
	"CaseCreationLimitExceeded":       newErrorCaseCreationLimitExceeded,
	"CaseIdNotFound":                  newErrorCaseIdNotFound,
	"DescribeAttachmentLimitExceeded": newErrorDescribeAttachmentLimitExceeded,
	"InternalServerError":             newErrorInternalServerError,
}
